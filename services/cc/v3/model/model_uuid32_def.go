package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Uuid32Def 实例ID。
type Uuid32Def struct {
}

func (o Uuid32Def) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Uuid32Def struct{}"
	}

	return strings.Join([]string{"Uuid32Def", string(data)}, " ")
}
