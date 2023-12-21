package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CharacterSetEnum struct {

	// 中文名。
	CnName *string `json:"cnName,omitempty"`

	// 编码。
	Code *string `json:"code,omitempty"`

	// 英文名。
	EnName *string `json:"enName,omitempty"`
}

func (o CharacterSetEnum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CharacterSetEnum struct{}"
	}

	return strings.Join([]string{"CharacterSetEnum", string(data)}, " ")
}
