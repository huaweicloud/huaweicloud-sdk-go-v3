package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubAppUpdateParam struct {

	// 子应用唯一标识
	Name string `json:"name"`

	// 子应用节点显示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 子应用描述
	Description *string `json:"description,omitempty"`
}

func (o SubAppUpdateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubAppUpdateParam struct{}"
	}

	return strings.Join([]string{"SubAppUpdateParam", string(data)}, " ")
}
