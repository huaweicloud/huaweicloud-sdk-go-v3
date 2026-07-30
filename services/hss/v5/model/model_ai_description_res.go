package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiDescriptionRes **参数解释**: 描述 **取值范围**: 字符长度0-256位
type AiDescriptionRes struct {
}

func (o AiDescriptionRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiDescriptionRes struct{}"
	}

	return strings.Join([]string{"AiDescriptionRes", string(data)}, " ")
}
