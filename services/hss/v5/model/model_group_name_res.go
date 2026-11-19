package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupNameRes **参数解释**: 服务器组名称 **取值范围**: 字符范围1-128位
type GroupNameRes struct {
}

func (o GroupNameRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupNameRes struct{}"
	}

	return strings.Join([]string{"GroupNameRes", string(data)}, " ")
}
