package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiIsDefaultRes **参数解释**: 是否是默认策略 **取值范围**: - false：否 - true：是
type AiIsDefaultRes struct {
}

func (o AiIsDefaultRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiIsDefaultRes struct{}"
	}

	return strings.Join([]string{"AiIsDefaultRes", string(data)}, " ")
}
