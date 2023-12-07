package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	HookUrl = "https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxxxxxxx" // test
)

func main() {
	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		data, _ := ioutil.ReadAll(c.Request.Body)
		err := Hook(string(data))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})
	})
	router.Run(":8002")
}

func Hook(data string) (err error) {
	log.Println("Hook start, data:", data)

	var notification Notification
	err = json.Unmarshal([]byte(data), &notification)
	if err != nil {
		log.Println("Hook Unmarshal err:", err)
		return
	}

	for _, v := range notification.Alerts {
		alert, _ := v.Labels["alertname"]
		instance, _ := v.Labels["instance"]
		des, _ := v.Annotations["description"]
		text := TempStr(alert, instance, des)
		_ = FeishuAlarmText(text)
	}

	return
}

func TempStr(alert, instance, des string) string {
	temp := `{"config":{"wide_screen_mode":true},"elements":[{"fields":[{"is_short":true,"text":{"content":"{{__alert__}}","tag":"lark_md"}},{"is_short":true,"text":{"content":"{{__instance__}}","tag":"lark_md"}}],"tag":"div"},{"tag":"div","text":{"content":"{{__text__}}","tag":"lark_md"}},{"tag":"hr"},{"elements":[{"content":"[来自 Prometheus](http://prometheus.staff.funlink-tech.com/)","tag":"lark_md"}],"tag":"note"}],"header":{"template":"red","title":{"content":"【Alert 报警】  {{__header__}}","tag":"plain_text"}}}`
	str := temp
	str = strings.ReplaceAll(str, "{{__header__}}", instance)
	str = strings.ReplaceAll(str, "{{__alert__}}", fmt.Sprintf("**类型:**  %s", alert))
	str = strings.ReplaceAll(str, "{{__instance__}}", fmt.Sprintf("**主机:**  [%s](http://grafana.staff.funlink-tech.com/d/9CWBz0bik/fu-wu-qi-xin-xi?orgId=1)", instance))
	str = strings.ReplaceAll(str, "{{__text__}}", fmt.Sprintf("**描述:**  %s", des))
	return fmt.Sprintf("{\"msg_type\":\"interactive\",\"card\":%v}", str)
}

func FeishuAlarmText(text string) (err error) {
	_, err = http.Post(HookUrl, "application/json", bytes.NewBufferString(text))
	if err != nil {
		log.Println("http err:", err)
		return
	}
	return
}