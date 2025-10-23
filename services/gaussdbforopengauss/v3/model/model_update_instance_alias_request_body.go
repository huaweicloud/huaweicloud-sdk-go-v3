package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceAliasRequestBody struct {

	// **参数解释**: 实例别名/备注。 **约束限制**: 不涉及。 **取值范围**: 允许中文，英文，数字及“-“、“_“、“.“，且长度为[1-64]个字符。 **默认取值**: 不涉及。
	Alias *string `json:"alias,omitempty"`
}

func (o UpdateInstanceAliasRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceAliasRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceAliasRequestBody", string(data)}, " ")
}
