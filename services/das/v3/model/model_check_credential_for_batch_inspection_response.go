package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCredentialForBatchInspectionResponse Response Object
type CheckCredentialForBatchInspectionResponse struct {

	// 测试结果
	CheckResult *bool `json:"check_result,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckCredentialForBatchInspectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCredentialForBatchInspectionResponse struct{}"
	}

	return strings.Join([]string{"CheckCredentialForBatchInspectionResponse", string(data)}, " ")
}
