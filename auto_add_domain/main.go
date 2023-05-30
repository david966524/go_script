package main

import (
	"domain/cf"
	"domain/fangneng"
	"flag"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/go-playground/validator"
)

var (
	domainFlag string //flag
	ipFlag     string //flag
)

var (
	key   string = "" //cf api key
	email string = "" //cf api email
)

// flag
func init() {
	flag.StringVar(&domainFlag, "domain", "", "输入添加的域名")
	flag.StringVar(&ipFlag, "ip", "", "输入回源ip地址")
}

// 真实服务
func start(domain string, ip string) {
	api, err := cloudflare.New(key, email) //cf api  client
	if err != nil {
		log.Fatal(err.Error())
	}
	//实例化 cf 对象
	cfobj := &cf.CloudFlare{
		Domain:   domain,
		OriginIp: ip,
		CnameUrl: "c0c5fe40.u.fn01.vip.",
		Api:      api,
	}
	cfobj.CreateDomain() //cf 添加域名方法
	cfobj.AddRecords()   //cf 添加解析方法
	//实例化 方能 对象
	fnobj := &fangneng.FangNeng{
		Name:        domain,
		UserPlanID:  62,
		ServerNames: []string{domain, "www." + domain, "mobile." + domain},
		Origins: []struct {
			Protocol  string `json:"protocol"`
			Host      string `json:"host"`
			PortRange string `json:"portRange"`
		}{
			{
				Protocol:  "http",
				Host:      ip,
				PortRange: "80",
			},
		},
		ListenPorts: []struct {
			Protocol  string `json:"protocol"`  //绑定端口协议，支持 http 或 https
			PortRange string `json:"portRange"` //绑定端口号
		}{
			{
				Protocol:  "http",
				PortRange: "80",
			},
			{
				Protocol:  "https",
				PortRange: "443",
			},
		},
		AutoCreateCert: true,
	}
	// 方能创建 网站 方法
	fnobj.CreateSite()
}

func main() {
	flag.Parse()
	//校验flag数据
	validate := validator.New()
	err := validate.Var(ipFlag, "ipv4")
	if err != nil {
		log.Println("ip格式错误  正确格式为:x.x.x.x")
		return
	}
	err1 := validate.Var(domainFlag, "fqdn")
	if err1 != nil {
		log.Println("域名格式错误 请检查！")
		return
	}
	log.Println("域名", domainFlag)
	log.Println("回源IP", ipFlag)
	//启动服务
	start(domainFlag, ipFlag)
}
