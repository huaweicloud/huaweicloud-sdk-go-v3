package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type DeploymentEdgecloud struct {

	// 边缘业务ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 边缘业务名称。
	Name *string `json:"name,omitempty" xml:"name"`

	Stacks *Stack `json:"stacks,omitempty" xml:"stacks"`

	// 边缘业务描述，最大支持255字节。
	Description *string `json:"description,omitempty" xml:"description"`

	Coverage *Coverage `json:"coverage,omitempty" xml:"coverage"`
}

func (o DeploymentEdgecloud) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentEdgecloud struct{}"
	}

	return strings.Join([]string{"DeploymentEdgecloud", string(data)}, " ")
}
