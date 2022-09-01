package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTestCaseDetailV2Request struct {

	// 用例编号，长度为[3-128]位字符
	TestcaseNumber string `json:"testcase_number" xml:"testcase_number"`
}

func (o ShowTestCaseDetailV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestCaseDetailV2Request struct{}"
	}

	return strings.Join([]string{"ShowTestCaseDetailV2Request", string(data)}, " ")
}
