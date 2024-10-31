package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceId 一种资源ID，32~36位的英文、数字、短横组合
type ResourceId struct {
}

func (o ResourceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceId struct{}"
	}

	return strings.Join([]string{"ResourceId", string(data)}, " ")
}
