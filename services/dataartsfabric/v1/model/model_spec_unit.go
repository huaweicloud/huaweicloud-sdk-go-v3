package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecUnit 单位
type SpecUnit struct {
}

func (o SpecUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecUnit struct{}"
	}

	return strings.Join([]string{"SpecUnit", string(data)}, " ")
}
