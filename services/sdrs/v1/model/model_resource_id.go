package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceId 保护实例Id
type ResourceId struct {

	// 资源ID
	Id string `json:"id"`
}

func (o ResourceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceId struct{}"
	}

	return strings.Join([]string{"ResourceId", string(data)}, " ")
}
