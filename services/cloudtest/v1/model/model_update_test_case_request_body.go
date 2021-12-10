package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新测试用例信息请求体
type UpdateTestCaseRequestBody struct {
	// 云测页面上显示的用例名称，长度为[3-128]位字符

	Name string `json:"name"`
	// 注册测试类型服务接口返回的服务id

	ServiceId int32 `json:"service_id"`
	// 测试用例等级，可选值为[0,1,2,3,4]，不填时云测默认为2

	RankId *string `json:"rank_id,omitempty"`
	// 用例编号，不填该值时云测会自动生成，长度为[3-128]位字符

	TestcaseNumber *string `json:"testcase_number,omitempty"`

	ExtendInfo *ExternalServiceBizCase `json:"extend_info,omitempty"`
}

func (o UpdateTestCaseRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseRequestBody", string(data)}, " ")
}
