package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTagInfo struct {

	// **参数解释**: 键 **约束限制**: key不能为空 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Key string `json:"key"`

	// **参数解释**: 值 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Value string `json:"value"`
}

func (o ResourceTagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagInfo struct{}"
	}

	return strings.Join([]string{"ResourceTagInfo", string(data)}, " ")
}
