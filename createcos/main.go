package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
)

func main() {

	secretID := ""
	secretKey := ""

	name := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)) //随机6为数字
	appid := "1318535612"                                                                        //腾讯云 appid
	bucket := fmt.Sprintf("%v-om-pro2-%v", name, appid)                                          //存储桶名
	urlstr := fmt.Sprintf("https://%v.cos.ap-hongkong.myqcloud.com", bucket)                     //存储桶url
	u, _ := url.Parse(urlstr)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,  //SecretID
			SecretKey: secretKey, //SecretKey
		},
	})

	// 创建存储桶
	opt := &cos.BucketPutOptions{
		XCosACL: "public-read-write",
	}
	resq, err := client.Bucket.Put(context.Background(), opt)
	if err != nil {
		panic(err)
	}
	log.Println(resq.Status)
	log.Println(urlstr)
	log.Printf("bucket Name: %v", bucket)

	// 设置配置文件的名字
	viper.SetConfigName("config")
	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 添加配置文件的路径，指定 config 目录下寻找
	viper.AddConfigPath(".")
	// 寻找配置文件并读取
	err1 := viper.ReadInConfig()
	if err1 != nil {
		panic(fmt.Errorf("fatal error config file: %w", err1))
	}
	appId := viper.Get("credential.tencent.appID")
	log.Println(appId)
	//修改appid
	viper.Set("credential.tencent.appID", appId)
	//修改bucket
	viper.Set("credential.tencent.bucket", bucket)
	//修改SecretID
	viper.Set("credential.tencent.SecretID", secretID)
	//修改SecretKey
	viper.Set("credential.tencent.SecretKey", secretKey)
	//保存配置文件
	err2 := viper.WriteConfig()
	if err2 != nil {
		log.Fatal("write config failed: ", err2)
	}

}
