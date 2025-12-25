package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentId 组件id.
type ComponentId struct {
}

func (o ComponentId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentId struct{}"
	}

	return strings.Join([]string{"ComponentId", string(data)}, " ")
}
