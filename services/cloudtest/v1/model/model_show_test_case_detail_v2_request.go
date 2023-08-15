package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestCaseDetailV2Request Request Object
type ShowTestCaseDetailV2Request struct {

	// 项目唯一标识，固定长度32位字符
	ProjectId string `json:"project_id"`

	// 用例编号，长度为[3-128]位字符
	TestcaseNumber string `json:"testcase_number"`
}

func (o ShowTestCaseDetailV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCaseDetailV2Request struct{}"
	}

	return strings.Join([]string{"ShowTestCaseDetailV2Request", string(data)}, " ")
}
