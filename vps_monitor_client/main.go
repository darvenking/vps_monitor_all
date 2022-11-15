package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-co-op/gocron"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"
	"vps_monitor_client/db"
	"vps_monitor_client/util"
)

func main() {

	timezone, _ := time.LoadLocation("Asia/Shanghai")
	schedule := gocron.NewScheduler(timezone)
	//schedule := gocron.NewScheduler(time.UTC)

	// 默认每30秒执行一次
	_, err1 := schedule.Every(5).Second().Do(func() {
		go handleSiteConfig()
	})
	// 默认每30秒执行一次
	_, err2 := schedule.Every(5).Second().Do(func() {
		go crawlStock()
	})
	if err1 != nil && err2 != nil {
		log.Println("定时任务启动失败...")
		return
	}
	schedule.StartBlocking()
}

// 加锁 防止在规定时间未执行完又重复执行
var handleUrlTaskRunningLock sync.Mutex

func handleSiteConfig() {
	if !handleUrlTaskRunningLock.TryLock() {
		log.Println("SiteConfig上一任务执行中，跳过当前定时任务...")
		return
	}
	defer handleUrlTaskRunningLock.Unlock()
	var sites []db.SiteConfig
	db.GetSiteConfigDB().Where("status = ?", 2).Find(&sites)
	for _, item := range sites {
		item := item
		//更新为已处理
		db.GetSiteConfigDB().Where("id = ?", item.ID).Update("status", 2)
		//处理网站内容
		if util.CheckUrl(item.URL) {
			//html, err := util.GetWebHtml(item.URL, item.Cookies)
			html, err := util.HttpGetWithHeader(item.URL, item.Cookies)
			if err != nil {
				continue
			}
			document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			if err != nil {
				continue
			}
			siteInfo := db.SiteInfo{
				URL:   item.URL,
				Price: document.Find(item.PriceFlag).Text(),
				Name:  document.Find(item.NameFlag).Text(),
			}
			//自动识别商家
			u, err := url.Parse(item.URL)
			if err != nil {
				log.Printf("url: 【%s】 --> 自动识别商家失败\n", item.URL)
			} else {
				var sellers []db.SellerInfo
				db.GetSellerInfoDB().Where("status = ?", 1).Find(&sellers)
				for _, v := range sellers {
					if strings.Contains(u.Host, v.SellerName) {
						siteInfo.SellerId = v.ID
						break
					}
				}
			}
			db.GetSiteInfoDB().Save(siteInfo)
			log.Printf("id: %d,url: 【%s】 --> 自动处理完成\n", item.ID, item.URL)
		}
	}

}

var handleSiteTaskRunningLock sync.Mutex

func crawlStock() {
	if !handleSiteTaskRunningLock.TryLock() {
		log.Println("CrawlStock上一任务执行中，跳过当前定时任务...")
		return
	}
	var sites []db.SiteInfo
	db.GetSiteInfoDB().Find(&sites) // 根据整型主键查找
	for _, item := range sites {
		item := item
		go crawlStockHandle(&item)
	}
	handleSiteTaskRunningLock.Unlock()
}

func crawlStockHandle(siteInfo *db.SiteInfo) {
	//result, err := util.GetWebHtml(siteInfo.URL, "")
	var config db.SiteConfig
	db.GetSiteConfigDB().First(&config, siteInfo.ConfigId)
	html, err := util.HttpGetWithHeader(siteInfo.URL, config.Cookies)
	if err != nil {
		return
	}
	res := false
	if config.NoStockFlag == "" {
		a := strings.Contains(html, "缺货")
		b := strings.Contains(html, "out of stock")
		res = !a && !b
	} else {
		res = !strings.Contains(html, config.NoStockFlag)
	}
	db.GetSiteInfoDB().Where("id = ?", siteInfo.ID).Update("stock", res)
	fmt.Printf("%s更新完成:%s,结果：%t", time.Now().Format("2006-01-02 15:04:05"), siteInfo.URL, res)
	fmt.Println()
}
