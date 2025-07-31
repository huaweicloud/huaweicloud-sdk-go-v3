package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessEuid **参数解释**: 进程有效用户ID **取值范围**: 最小值0，最大值9223372036854775807
type ProcessEuid struct {
}

func (o ProcessEuid) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessEuid struct{}"
	}

	return strings.Join([]string{"ProcessEuid", string(data)}, " ")
}
