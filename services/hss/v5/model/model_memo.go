package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Memo **参数解释** 备注信息 **取值范围** 字符长度0-512位
type Memo struct {
}

func (o Memo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Memo struct{}"
	}

	return strings.Join([]string{"Memo", string(data)}, " ")
}
