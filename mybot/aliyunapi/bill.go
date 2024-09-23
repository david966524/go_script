// This file is auto-generated, don't edit it. Thanks.
package aliyunapi

import (
	"strconv"
	"time"

	bssopenapi20171214 "github.com/alibabacloud-go/bssopenapi-20171214/v5/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// Description:
//
// 使用AK&SK初始化账号Client
//
// @return Client
//
// @throws Exception
func CreateClient() (_result *bssopenapi20171214.Client, _err error) {
	// 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
	// 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(""),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(""),
	}
	// Endpoint 请参考 https://api.aliyun.com/product/BssOpenApi
	config.Endpoint = tea.String("business.ap-southeast-1.aliyuncs.com")
	_result = &bssopenapi20171214.Client{}
	_result, _err = bssopenapi20171214.NewClient(config)
	return _result, _err
}

func AccoutBanlance() string {
	runtime := &util.RuntimeOptions{}
	cli, _ := CreateClient()
	result, err := cli.QueryAccountBalanceWithOptions(runtime)
	if err != nil {
		return err.Error()
	}
	return "账户余额" + *result.Body.Data.AvailableAmount + *result.Body.Data.Currency
}

func DailyBill() string {
	// 获取当前时间
	now := time.Now()

	// 计算昨天的时间
	yesterday := now.AddDate(0, 0, -1)

	// 格式化日期为所需格式，例如 "2024-08-25"
	yesterdayDate := yesterday.Format("2006-01-02")
	monthDate := yesterday.Format("2006-01")
	queryAccountBillRequest := &bssopenapi20171214.QueryAccountBillRequest{
		BillingDate:  tea.String(yesterdayDate),
		BillingCycle: tea.String(monthDate),
		Granularity:  tea.String("DAILY"),
	}
	cli, err := CreateClient()
	if err != nil {
		return err.Error()
	}
	runtime := &util.RuntimeOptions{}
	result, err := cli.QueryAccountBillWithOptions(queryAccountBillRequest, runtime)
	if err != nil {
		return err.Error()
	}
	return yesterdayDate + "----" + strconv.FormatFloat(float64(*result.Body.Data.Items.Item[0].PretaxAmount), 'f', 2, 32) + *result.Body.Data.Items.Item[0].Currency
}

func main() {
	DailyBill()

}
