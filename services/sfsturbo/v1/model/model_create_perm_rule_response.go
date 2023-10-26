package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePermRuleResponse Response Object
type CreatePermRuleResponse struct {

	// 权限规格信息
	Rules          *[]OnePermRuleResponseInfo `json:"rules,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o CreatePermRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePermRuleResponse struct{}"
	}

	return strings.Join([]string{"CreatePermRuleResponse", string(data)}, " ")
}
