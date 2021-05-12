package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetSchUrl 查询每日库存 && 能否预约等
var GetSchUrl = "https://www.wechat.lgjkzx.com/mps-area-version/appointment/getSchForPost"

var MyHttpClient = NewHttpClient(nil)

func getSchInfo() (result *Result, err error) {
	headers := http.Header{
		"Connection":       []string{"keep-alive"},
		"Cache-Control":    []string{"max-age=0"},
		"Content-Length":   []string{"53"},
		"Accept":           []string{"application/json, text/javascript, */*;q=0.01"},
		"Origin":           []string{"https://www.wechat.lgjkzx.com"},
		"X-Requested-With": []string{"XMLHttpRequest"},
		"User-Agent":       []string{"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36 QBCore/4.0.1326.400 QQBrowser/9.0.2524.400 Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2875.116 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63010200)"},
		"Content-Type":     []string{"application/x-www-form-urlencoded; charset=UTF-8"},
		"Referer":          []string{"https://www.wechat.lgjkzx.com/mps-area-version/appointment/getDoctorInfoAndSrc?param=pfWQja1m4KzmtQ6GsxRU763Sg82GySEaBlw0oyjjmZTlvR5I1yOQ8WjKr59ife4OE66tzKl0gJCsoCmyFth5%2BH6kpqVbz5hjm4vBYFmPMn%2Fq4AWhhLBDFuGjhIjXX8KPt7RBaydQv3obeIO0KyPxWUf6Y8mhn3CYzahnSHJDr6Y%2F0TA4%2BUx3usQAHQeIbrTP"},
		"Accept-Encoding":  []string{"gzip, deflate"},
		"Accept-Language":  []string{"zh-CN,zh;q=0.8,en-US;q=0.6,en;q=0.5;q=0.4"},
		"Cookie":           []string{"SESSION=429d3520-3c1f-489b-a88f-213db14ea785"},
	}
	var cookies []http.Cookie
	body := bytes.NewReader([]byte("hospitalId=62&deptId=7381&doctorId=35734&numberDays=1"))
	response, err := MyHttpClient.doRequest(context.Background(), http.MethodPost, GetSchUrl, headers, cookies, body)
	if err != nil {
		fmt.Printf("do request failed, err:%v\n", err)
		return
	}
	if response.Status != http.StatusOK {
		fmt.Printf("请求失败，响应状态码：%d\n", response.Status)
		return
	}
	fmt.Printf("获取到的数据：%s\n", string(response.Body))
	err = json.Unmarshal(response.Body, &result)
	return
}

func main() {
	result, err := getSchInfo()
	if err != nil {
		fmt.Printf("获取库存信息失败：%v\n", err)
	}
	fmt.Println(result)
}
