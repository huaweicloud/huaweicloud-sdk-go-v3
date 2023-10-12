package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Enabled 开关
type Enabled struct {
}

func (o Enabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Enabled struct{}"
	}

	return strings.Join([]string{"Enabled", string(data)}, " ")
}
