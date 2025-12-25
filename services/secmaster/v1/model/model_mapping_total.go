package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MappingTotal 总数
type MappingTotal struct {
}

func (o MappingTotal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MappingTotal struct{}"
	}

	return strings.Join([]string{"MappingTotal", string(data)}, " ")
}
