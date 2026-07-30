package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssociatedResourceRulesResponse Response Object
type UpdateAssociatedResourceRulesResponse struct {

	// 本次操作的规则信息
	Rules *[]AssociatedResourceRule `json:"rules,omitempty"`

	// 操作的错误信息
	Errors         *[]ErrorInfo `json:"errors,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateAssociatedResourceRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssociatedResourceRulesResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssociatedResourceRulesResponse", string(data)}, " ")
}
