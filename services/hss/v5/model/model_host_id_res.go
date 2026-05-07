package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostIdRes **参数解释**: 服务器的唯一标识ID **取值范围**: 字符长度1-64位
type HostIdRes struct {
}

func (o HostIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostIdRes struct{}"
	}

	return strings.Join([]string{"HostIdRes", string(data)}, " ")
}
