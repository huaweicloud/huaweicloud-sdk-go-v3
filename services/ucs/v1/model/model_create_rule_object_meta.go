package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRuleObjectMeta struct {

	// 权限策略名称
	Name string `json:"name"`
}

func (o CreateRuleObjectMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleObjectMeta struct{}"
	}

	return strings.Join([]string{"CreateRuleObjectMeta", string(data)}, " ")
}
