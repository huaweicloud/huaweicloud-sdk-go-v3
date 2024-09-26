package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Uuid64Def 实例ID。
type Uuid64Def struct {
}

func (o Uuid64Def) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Uuid64Def struct{}"
	}

	return strings.Join([]string{"Uuid64Def", string(data)}, " ")
}
