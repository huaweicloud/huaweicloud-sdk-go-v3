package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NameDef 实例名称。
type NameDef struct {
}

func (o NameDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NameDef struct{}"
	}

	return strings.Join([]string{"NameDef", string(data)}, " ")
}
