package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DebugLiveDataApiV2Response struct {
	// 测试返回的状态码

	TestStatusCode *string `json:"test_status_code,omitempty"`
	// 测试的请求内容

	TestRequest *string `json:"test_request,omitempty"`
	// 测试耗时

	TestUsedTime *int32 `json:"test_used_time,omitempty"`
	// 测试者的项目编号

	TestOperator *string `json:"test_operator,omitempty"`
	// 测试的响应内容

	TestResponse *string `json:"test_response,omitempty"`
	// 测试的请求方法

	TestMethod *string `json:"test_method,omitempty"`
	// 测试编号

	TestId *int32 `json:"test_id,omitempty"`
	// 测试时间

	TestDate *sdktime.SdkTime `json:"test_date,omitempty"`
	// 后端API编号

	LdApiId *string `json:"ld_api_id,omitempty"`
	// 本次测试日志列表

	DebugLog       *[]string `json:"debug_log,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DebugLiveDataApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugLiveDataApiV2Response struct{}"
	}

	return strings.Join([]string{"DebugLiveDataApiV2Response", string(data)}, " ")
}
