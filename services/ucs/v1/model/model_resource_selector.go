package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceSelector struct {

	// 目标资源的API版本
	ApiVersion *string `json:"apiVersion,omitempty"`

	// 目标资源的类型
	Kind *string `json:"kind,omitempty"`

	// 目标资源的名称
	Namespace *string `json:"namespace,omitempty"`

	// 目标资源所在命名空间
	Name *string `json:"name,omitempty"`
}

func (o ResourceSelector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSelector struct{}"
	}

	return strings.Join([]string{"ResourceSelector", string(data)}, " ")
}
