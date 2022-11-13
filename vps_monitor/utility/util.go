package utility

import (
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

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

func String2bytes(s string) []byte {
	tmp1 := (*[2]uintptr)(unsafe.Pointer(&s))
	tmp2 := [3]uintptr{tmp1[0], tmp1[1], tmp1[1]}
	return *(*[]byte)(unsafe.Pointer(&tmp2))
}

func Bytes2string(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
