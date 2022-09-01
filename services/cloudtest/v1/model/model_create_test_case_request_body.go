package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建测试用例请求体
type CreateTestCaseRequestBody struct {

	// 云测页面上显示的用例名称，长度为[3-128]位字符
	Name string `json:"name" xml:"name"`

	// 该值由注册接口返回,取值范围为10-9999
	ServiceId int32 `json:"service_id" xml:"service_id"`

	// 测试用例等级，可选值为[0,1,2,3,4]，不填时云测默认为2
	RankId *string `json:"rank_id,omitempty" xml:"rank_id"`

	// 用例编号，不填该值时云测会自动生成，长度为[3-128]位字符
	TestcaseNumber *string `json:"testcase_number,omitempty" xml:"testcase_number"`

	ExtendInfo *ExternalServiceCaseInfo `json:"extend_info,omitempty" xml:"extend_info"`
}

func (o CreateTestCaseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTestCaseRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTestCaseRequestBody", string(data)}, " ")
}
