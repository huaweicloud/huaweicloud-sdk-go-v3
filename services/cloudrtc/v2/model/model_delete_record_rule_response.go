package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecordRuleResponse Response Object
type DeleteRecordRuleResponse struct {
	XRequestId     *string `json:"X-request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRecordRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecordRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecordRuleResponse", string(data)}, " ")
}
