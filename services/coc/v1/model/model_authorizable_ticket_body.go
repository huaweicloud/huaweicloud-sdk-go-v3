package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuthorizableTicketBody struct {

	// 可授权单类型
	Type *string `json:"type,omitempty"`

	// scope ID，一般为region id
	ScopeId *string `json:"scope_id,omitempty"`

	// 目标 ID，一般为应用id
	TargetId *string `json:"target_id,omitempty"`
}

func (o AuthorizableTicketBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizableTicketBody struct{}"
	}

	return strings.Join([]string{"AuthorizableTicketBody", string(data)}, " ")
}
