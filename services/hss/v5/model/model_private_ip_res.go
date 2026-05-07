package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrivateIpRes **参数解释**: 服务器私有IP **取值范围**: 字符长度1-128位
type PrivateIpRes struct {
}

func (o PrivateIpRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateIpRes struct{}"
	}

	return strings.Join([]string{"PrivateIpRes", string(data)}, " ")
}
