package util

import (
	"context"
	"github.com/chromedp/chromedp"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"unsafe"
	"vps_monitor_client/db"
)

func GetSiteInfo(item *db.SitePre) (*db.SiteInfo, error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		// debug使用
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	c, _ := chromedp.NewExecAllocator(context.Background(), options...)
	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)
	//创建一个上下文，超时时间为40s
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()
	//var htmlContent string
	var name, price string
	err := chromedp.Run(timeoutCtx,
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

func String2bytes(s string) []byte {
	tmp1 := (*[2]uintptr)(unsafe.Pointer(&s))
	tmp2 := [3]uintptr{tmp1[0], tmp1[1], tmp1[1]}
	return *(*[]byte)(unsafe.Pointer(&tmp2))
}

func Bytes2string(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

/*
*判断url地址是否合规
 */
func CheckUrl(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

func HttpGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		if Body.Close() != nil {
			return
		}
	}(resp.Body)
	//循环读取网页数据
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			break
		}
		//累加每一次循环读取到的数据
		result += string(buf[:n])
	}
	return
}
