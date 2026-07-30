package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssociatedResourceRuleResponse Response Object
type DeleteAssociatedResourceRuleResponse struct {

	// 操作失败的错误信息
	Errors         *[]ErrorInfo `json:"errors,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteAssociatedResourceRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssociatedResourceRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssociatedResourceRuleResponse", string(data)}, " ")
}
