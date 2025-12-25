package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentName 组件名称
type ComponentName struct {
}

func (o ComponentName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentName struct{}"
	}

	return strings.Join([]string{"ComponentName", string(data)}, " ")
}
