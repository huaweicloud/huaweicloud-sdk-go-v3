package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRulesResponse Response Object
type ShowRulesResponse struct {

	// 规则列表
	Rules *[]ResponseRule `json:"rules,omitempty"`

	// 规则总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRulesResponse struct{}"
	}

	return strings.Join([]string{"ShowRulesResponse", string(data)}, " ")
}
