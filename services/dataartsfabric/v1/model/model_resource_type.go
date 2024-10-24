package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceType 资源类型
type ResourceType struct {
}

func (o ResourceType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceType struct{}"
	}

	return strings.Join([]string{"ResourceType", string(data)}, " ")
}
