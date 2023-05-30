package cf

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

type CloudFlare struct {
	Domain   string
	OriginIp string
	CnameUrl string
	Api      *cloudflare.API
	Id       string
}

func (cf *CloudFlare) CreateDomain() {
	zoneParams := cloudflare.ZoneCreateParams{
		Name:      cf.Domain,
		JumpStart: false,
	}
	log.Println("====================开始向cf 添加域名======================")
	zone, err := cf.Api.CreateZone(context.Background(), zoneParams.Name, zoneParams.JumpStart, cloudflare.Account{}, "")

	if err != nil {
		log.Fatal(err)
	}
	cf.Id = zone.ID
	log.Println("当前zone 状态为：" + zone.Status)
}

func (cf *CloudFlare) AddRecords() {
	domainMap := make(map[string]string)
	domainMap[fmt.Sprintf("%v.%v", "www", cf.Domain)] = "CNAME"
	domainMap[fmt.Sprintf("%v.%v", "mobile", cf.Domain)] = "CNAME"
	domainMap[cf.Domain] = "CNAME"
	domainMap[fmt.Sprintf("%v.%v", "admin", cf.Domain)] = "A"
	proxied := true
	log.Println("====================开始添加解析======================")

	for k, v := range domainMap {
		if v == "CNAME" {
			proxied = false
			dnsRecordParams := cloudflare.CreateDNSRecordParams{
				Type:    v,
				Name:    k,
				Content: cf.CnameUrl,
				Proxied: &proxied,
				TTL:     300,
			}
			record, err := cf.Api.CreateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(cf.Id), dnsRecordParams)
			log.Println(fmt.Sprintf("添加解析 %v  类型 %v  解析值 %v TTL值 %v", record.Name, record.Type, record.Content, record.TTL))
			if err != nil {
				log.Println(err.Error())
			}
		} else {
			proxied = true
			dnsRecordParams := cloudflare.CreateDNSRecordParams{
				Type:    v,
				Name:    k,
				Content: cf.OriginIp,
				Proxied: &proxied,
				TTL:     300,
			}
			record, err := cf.Api.CreateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(cf.Id), dnsRecordParams)
			log.Println(fmt.Sprintf("添加解析 %v  类型 %v  解析值 %v TTL值 %v", record.Name, record.Type, record.Content, record.TTL))
			if err != nil {
				log.Println(err.Error())
			}
		}

	}

}
