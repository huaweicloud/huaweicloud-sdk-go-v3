package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LdApiTestHistoryInfoV2 struct {

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
}

func (o LdApiTestHistoryInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiTestHistoryInfoV2 struct{}"
	}

	return strings.Join([]string{"LdApiTestHistoryInfoV2", string(data)}, " ")
}
