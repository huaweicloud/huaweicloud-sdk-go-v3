package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Value 阈值
type Value struct {
}

func (o Value) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Value struct{}"
	}

	return strings.Join([]string{"Value", string(data)}, " ")
}
