package util

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"net"
	"time"
	"vps_monitor_client/db"
)

func GetSiteInfo(item *db.SiteConfig) (*db.SiteInfo, error) {
	chromeCtx, cancel := context.WithTimeout(GetChromeCtx(false), 40*time.Second)
	defer cancel()
	//var htmlContent string
	var name, price string
	err := chromedp.Run(chromeCtx,
		chromedp.Navigate(item.URL),
		chromedp.TextContent(item.NameFlag, &name),
		chromedp.TextContent(item.PriceFlag, &price),
	)
	if err != nil {
		log.Printf("Run err : %v\n", err)
		return nil, err
	}
	d := &db.SiteInfo{
		URL:         item.URL,
		Name:        name,
		Price:       price,
		NoStockFlag: item.NoStockFlag,
	}
	return d, nil
}

func GetWebHtml(url string) (string, error) {
	chromeCtx, cancel := context.WithTimeout(GetChromeCtx(false), 40*time.Second)
	defer cancel()
	var htmlContent string
	err := chromedp.Run(chromeCtx,
		chromedp.Navigate(url),
		chromedp.OuterHTML(`/html/body`, &htmlContent),
	)
	if err != nil {
		log.Printf("Run err : %v\n", err)
		return "", err
	}
	return htmlContent, nil
}

// chromeCtx 使用一个实例
var chromeCtx context.Context

func GetChromeCtx(focus bool) context.Context {
	if chromeCtx == nil || focus {
		//初始化参数，先传一个空的数据
		allocOpts := chromedp.DefaultExecAllocatorOptions[:]
		allocOpts = append(allocOpts,
			chromedp.DisableGPU,
			chromedp.Flag("headless", true),
			chromedp.Flag("blink-settings", "imagesEnabled=false"),
			chromedp.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36`),
			chromedp.Flag("accept-language", `zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6`),
		)
		var c context.Context
		if checkChromePort() {
			// 不知道为何，不能直接使用 NewExecAllocator ，因此增加 使用 ws://127.0.0.1:9222/ 来调用
			c, _ = chromedp.NewRemoteAllocator(context.Background(), "ws://127.0.0.1:9222/")
		} else {
			c, _ = chromedp.NewExecAllocator(context.Background(), allocOpts...)
		}
		chromeCtx, _ = chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
		// 执行一个空task, 用提前创建Chrome实例
		chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)
	}

	return chromeCtx
}

// 检查是否有9222端口，来判断是否运行在linux上
func checkChromePort() bool {
	addr := net.JoinHostPort("", "9222")
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
