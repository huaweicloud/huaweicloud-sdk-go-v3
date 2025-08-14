package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrivateIp **参数解释**： 服务器私有IP **取值范围**： 字符长度1-128位
type PrivateIp struct {
}

func (o PrivateIp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateIp struct{}"
	}

	return strings.Join([]string{"PrivateIp", string(data)}, " ")
}
