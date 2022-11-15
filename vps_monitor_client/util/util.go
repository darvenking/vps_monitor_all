package util

import (
	"io"
	"log"
	"net/http"
	"strings"
	"unsafe"
)

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

func HttpGetWithHeader(url string, cookies string) (result string, err error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("err")
	}
	// 添加请求头
	if cookies != "" {
		req.Header.Add("cookie", cookies)
	}
	req.Header.Add("user-agent", GetCfgStr("user-agent"))
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	result = string(b)
	return
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
