package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UseRuleVo struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`
}

func (o UseRuleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UseRuleVo struct{}"
	}

	return strings.Join([]string{"UseRuleVo", string(data)}, " ")
}
