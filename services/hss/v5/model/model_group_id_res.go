package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupIdRes **参数解释**: 服务器组ID **取值范围**: 字符范围1-64位
type GroupIdRes struct {
}

func (o GroupIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupIdRes struct{}"
	}

	return strings.Join([]string{"GroupIdRes", string(data)}, " ")
}
