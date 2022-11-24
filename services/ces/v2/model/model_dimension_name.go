package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源维度，必须以字母开头，多维度用\",\"分割，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32
type DimensionName struct {
}

func (o DimensionName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionName struct{}"
	}

	return strings.Join([]string{"DimensionName", string(data)}, " ")
}
