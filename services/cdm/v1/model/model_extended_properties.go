package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtendedProperties 集群扩展信息
type ExtendedProperties struct {

	// 工作空间ID。
	WorkSpaceId *string `json:"workSpaceId,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resourceId,omitempty"`

	// 是否是试用集群。
	Trial *string `json:"trial,omitempty"`
}

func (o ExtendedProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendedProperties struct{}"
	}

	return strings.Join([]string{"ExtendedProperties", string(data)}, " ")
}
