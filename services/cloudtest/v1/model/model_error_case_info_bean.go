package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ErrorCaseInfoBean struct {

	// 失败错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 失败错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// 失败资源信息
	TestcaseId *string `json:"testcase_id,omitempty" xml:"testcase_id"`
}

func (o ErrorCaseInfoBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorCaseInfoBean struct{}"
	}

	return strings.Join([]string{"ErrorCaseInfoBean", string(data)}, " ")
}
