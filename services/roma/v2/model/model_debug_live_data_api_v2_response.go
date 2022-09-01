package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DebugLiveDataApiV2Response struct {

	// 测试返回的状态码
	TestStatusCode *string `json:"test_status_code,omitempty" xml:"test_status_code"`

	// 测试的请求内容
	TestRequest *string `json:"test_request,omitempty" xml:"test_request"`

	// 测试耗时
	TestUsedTime *int32 `json:"test_used_time,omitempty" xml:"test_used_time"`

	// 测试者的项目编号
	TestOperator *string `json:"test_operator,omitempty" xml:"test_operator"`

	// 测试的响应内容
	TestResponse *string `json:"test_response,omitempty" xml:"test_response"`

	// 测试的请求方法
	TestMethod *string `json:"test_method,omitempty" xml:"test_method"`

	// 测试编号
	TestId *int32 `json:"test_id,omitempty" xml:"test_id"`

	// 测试时间
	TestDate *sdktime.SdkTime `json:"test_date,omitempty" xml:"test_date"`

	// 后端API编号
	LdApiId *string `json:"ld_api_id,omitempty" xml:"ld_api_id"`

	// 本次测试日志列表
	DebugLog       *[]string `json:"debug_log,omitempty" xml:"debug_log"`
	HttpStatusCode int       `json:"-"`
}

func (o DebugLiveDataApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugLiveDataApiV2Response struct{}"
	}

	return strings.Join([]string{"DebugLiveDataApiV2Response", string(data)}, " ")
}
