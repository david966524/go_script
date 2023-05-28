package fangneng

type RespMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
	} `json:"data"`
}

type FangNeng struct {
	Name        string   `json:"name"`        //网站名称，必填
	UserPlanID  int      `json:"userPlanId"`  //用户套餐ID，可由套餐列表接口获取，必填
	ServerNames []string `json:"serverNames"` //网站域名，可多域名，支持泛域名，必填
	Origins     []struct {
		Protocol  string `json:"protocol"`  //源站协议，支持 http 或 https
		Host      string `json:"host"`      //源站地址，可输入IP或域名
		PortRange string `json:"portRange"` //源站端口号
	} `json:"origins"`
	AutoCreateCert bool `json:"autoCreateCert"` //自动申请https证书，启用后自动申请证书并与网站绑定
}

func (fn *FangNeng) CreateSite() {

}
