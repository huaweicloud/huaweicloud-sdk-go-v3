package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源分组ID，监控范围为资源分组时必传
type ResourceGroupId struct {
}

func (o ResourceGroupId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroupId struct{}"
	}

	return strings.Join([]string{"ResourceGroupId", string(data)}, " ")
}
