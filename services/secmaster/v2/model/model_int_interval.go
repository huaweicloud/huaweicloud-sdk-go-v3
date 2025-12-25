package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IntInterval 整型时长
type IntInterval struct {
}

func (o IntInterval) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntInterval struct{}"
	}

	return strings.Join([]string{"IntInterval", string(data)}, " ")
}
