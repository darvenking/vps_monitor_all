package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"
	"vps_monitor_client/db"
	"vps_monitor_client/util"
)

func main() {
	go handleSiteConfig()
	go crawlStock()
	select {}
}

func handleSiteConfig() {
	for {
		var sites []db.SiteConfig
		db.GetSiteConfigDB().Where("status = ?", 2).Find(&sites)
		for _, item := range sites {
			item := item
			//更新为已处理
			db.GetSiteConfigDB().Where("id = ?", item.ID).Update("status", 2)
			//处理网站内容
			if util.CheckUrl(item.URL) {
				siteInfo, err := util.GetSiteInfo(&item)
				if err != nil {
					continue
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
		time.Sleep(300 * time.Second)
	}
}

func crawlStock() {
	for {
		var sites []db.SiteInfo
		db.GetSiteInfoDB().Find(&sites) // 根据整型主键查找
		for _, item := range sites {
			item := item
			go handle(&item)
		}
		time.Sleep(30 * time.Second)
	}
}

func handle(siteInfo *db.SiteInfo) {
	result, err := util.GetWebHtml(siteInfo.URL)
	if err != nil {
		return
	}
	res := false
	if siteInfo.NoStockFlag == "" {
		a := strings.Contains(result, "缺货")
		b := strings.Contains(result, "out of stock")
		res = !a && !b
	} else {
		res = !strings.Contains(result, siteInfo.NoStockFlag)
	}
	db.GetSiteInfoDB().Where("id = ?", siteInfo.ID).Update("stock", res)
	fmt.Printf("%s更新完成:%s,结果：%t", time.Now().Format("2006-01-02 15:04:05"), siteInfo.URL, res)
	fmt.Println()
}
