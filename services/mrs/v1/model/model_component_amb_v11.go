package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentAmbV11 struct {

	// 组件名称
	ComponentName string `json:"component_name"`
}

func (o ComponentAmbV11) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentAmbV11 struct{}"
	}

	return strings.Join([]string{"ComponentAmbV11", string(data)}, " ")
}
