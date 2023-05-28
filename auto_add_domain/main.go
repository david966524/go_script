package main

import (
	"flag"
	"log"

	"github.com/go-playground/validator"
)

var (
	domainFlag string
	ipFlag     string
)

func init() {

	flag.StringVar(&domainFlag, "domain", "", "输入添加的域名")
	flag.StringVar(&ipFlag, "ip", "", "输入回源ip地址")
}

func main() {
	flag.Parse()

	validate := validator.New()

	err := validate.Var(ipFlag, "ipv4")
	if err != nil {
		log.Println("ip格式错误  正确格式为：x.x.x.x")
		return
	}
	err1 := validate.Var(domainFlag, "fqdn")
	if err1 != nil {
		log.Println("域名格式错误 请检查！")
		return
	}
	log.Println("域名", domainFlag)
	log.Println("回源IP", ipFlag)
}
