package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UuidDef 资源ID标识符。
type UuidDef struct {
}

func (o UuidDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UuidDef struct{}"
	}

	return strings.Join([]string{"UuidDef", string(data)}, " ")
}
