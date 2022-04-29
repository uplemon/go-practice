package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func CookieDemo() {
	http.HandleFunc("/", testCookieHandler)
	_ = http.ListenAndServe(":8085", nil)
}

func testCookieHandler(w http.ResponseWriter, r *http.Request) {
	// 从Request中获取cookie
	c, err := r.Cookie("test_cookie")
	fmt.Printf("cookie:%#v, err:%v\n", c, err)
	// 设置cookie
	cookie := &http.Cookie{
		Name:   "test_cookie",
		Value:  "Go-Web" + strconv.FormatInt(time.Now().UnixNano(), 10),
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}
	/**
	 * 应在具体数据返回之前设置Cookie，否则设置不成功
	 * http.SetCookie
	 */
	http.SetCookie(w, cookie)
	w.Write([]byte("hello world."))
}

func cookieDesc() {
	// Cookie的结构体定义：http.Cookie{}

	// 设置cookie
	// func SetCookie(w ResponseWriter, cookie *Cookie)

	/* net/http包中的Request对象3个处理cookie的方法 */
	// func (r *Request) Cookies() []*Cookie
	// func (r *Request) Cookie(name string) (*Cookie, error)
	// func (r *Request) AddCookie(c *Cookie)
}

func GetUrlContent(method, urlVal, data string) {
	var client = &http.Client{}
	var req *http.Request
	if data == "" {
		urlArr := strings.Split(urlVal, "?")
		if len(urlArr) == 2 {
			urlVal = urlArr[0] + "?" + url.PathEscape(urlArr[1])
		}
		req, _ = http.NewRequest(method, urlVal, nil)
	} else {
		req, _ = http.NewRequest(method, urlVal, strings.NewReader(data))
	}
	// 添加cookie
	cookie := &http.Cookie{
		Name:  "x-csrf-token",
		Value: "abc2def1gfl6kds1kfg0hnv3laf1kdl9sfl2s3fd6mnn0d",
	}
	req.AddCookie(cookie)

	// 添加header
	req.Header.Add("x-request-id", "f1gfl6kds1kfg0hn")

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	// 读取数据
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
