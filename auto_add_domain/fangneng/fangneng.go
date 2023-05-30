package fangneng

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type RespMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
	} `json:"data"`
}

type FangNeng struct {
	Name        string   `json:"name"`        //网站名称，必填
	UserPlanID  int      `json:"userPlanId"`  //用户套餐ID，可由套餐列表接口获取，必填   "id":62,"name":"100435_kjzd8341-package-100435_kjzd8341-package-user"
	ServerNames []string `json:"serverNames"` //网站域名，可多域名，支持泛域名，必填
	Origins     []struct {
		Protocol  string `json:"protocol"`  //源站协议，支持 http 或 https
		Host      string `json:"host"`      //源站地址，可输入IP或域名
		PortRange string `json:"portRange"` //源站端口号
	} `json:"origins"`
	ListenPorts []struct {
		Protocol  string `json:"protocol"`  //绑定端口协议，支持 http 或 https
		PortRange string `json:"portRange"` //绑定端口号
	} `json:"listenPorts"`
	AutoCreateCert bool `json:"autoCreateCert"` //自动申请https证书，启用后自动申请证书并与网站绑定
}

func (fn *FangNeng) CreateSite() {
	log.Println(fmt.Sprintf(`
	服务名称 %v
	绑定域名 %v
	源站地址 %v
	申请证书 %v
	`, fn.Name, fn.ServerNames, fn.Origins, fn.AutoCreateCert))
	var respMsg RespMsg
	headersmap := map[string]string{
		"Content-Type": "application/json",
		"X-API-Key":    "", //方能api key
		"X-API-Secret": "", //方能api ID
	}
	url := "https://api1.funnull.io" + "/api/server"
	log.Println("==============向方能添加 域名==============")
	client := resty.New()
	resp, err := client.R().EnableTrace().
		SetHeaders(headersmap).SetBody(fn).SetResult(&respMsg).Post(url)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("========方能返回结果========")
	log.Println(resp.String())
	log.Println("==========生成域名链接======")
	for _, v := range fn.ServerNames {
		fmt.Printf("https://%v \n", v)
	}
}
