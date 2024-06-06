package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNoticeRuleResponse Response Object
type DeleteNoticeRuleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteNoticeRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNoticeRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteNoticeRuleResponse", string(data)}, " ")
}
