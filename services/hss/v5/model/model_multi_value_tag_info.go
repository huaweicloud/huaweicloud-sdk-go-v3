package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiValueTagInfo struct {

	// **参数解释**: 标签键。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位。 **默认取值**: 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释**: 标签值列表。 **约束限制**: 不涉及 **取值范围**: 每个值字符长度1-256位。 **默认取值**: 不涉及
	Values *[]string `json:"values,omitempty"`
}

func (o MultiValueTagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiValueTagInfo struct{}"
	}

	return strings.Join([]string{"MultiValueTagInfo", string(data)}, " ")
}
