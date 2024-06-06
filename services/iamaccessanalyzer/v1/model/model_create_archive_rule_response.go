package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateArchiveRuleResponse Response Object
type CreateArchiveRuleResponse struct {

	// 存档规则的唯一标识符。
	Id *string `json:"id,omitempty"`

	// 存档规则的唯一资源标识符。
	Urn            *string `json:"urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateArchiveRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateArchiveRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateArchiveRuleResponse", string(data)}, " ")
}
