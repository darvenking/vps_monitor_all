package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"vps_monitor_client/db"
)

func main() {

	var sites []db.SiteInfo
	db.GetSiteInfoDB().Find(&sites) // 根据整型主键查找
	for _, item := range sites {
		item := item
		go handle(&item)
	}

	select {}

}

func handle(siteInfo *db.SiteInfo) {
	for {
		result, err := HttpGet(siteInfo.URL)
		if err != nil {
			return
		}
		db.GetSiteInfoDB().Where("id = ?", siteInfo.ID).Update("stock", !strings.Contains(result, "out of stock"))
		time.Sleep(30 * time.Second)
	}

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
			fmt.Println("读取完成")
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
