package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MappingId 映射id
type MappingId struct {
}

func (o MappingId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MappingId struct{}"
	}

	return strings.Join([]string{"MappingId", string(data)}, " ")
}
