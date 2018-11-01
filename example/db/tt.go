package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("st")

	str := "http://baidu.com?q=http%3A%2F%2Fweibo.cn"
	vls,_ := url.ParseQuery(str)
	fmt.Println(vls)

	v, err := url.ParseRequestURI("http://www.baidu.com/s?q2=http%3A%2F%2Fweibo.cn&wd=%E5%BE%AE%E5%BA%A6%E7%BD%91%E7%BB%9C&rsv_spt=1&issp=1&rsv_bp=0&ie=utf-8&tn=baiduhome_pg&rsv_sug3=2&rsv_sug4=172&rsv_sug1=1")
	if err != nil {
		fmt.Println(err)
	}
	u := v.RawQuery
	//parsequery传入的必须是参数，也就是url里边的RawQuery的值 就是url?之后的path
	fmt.Println(url.ParseQuery(u))
	//这里url.Query()直接就解析成map了
	fmt.Println(v.Query())

}
