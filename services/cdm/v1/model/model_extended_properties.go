package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群扩展信息
type ExtendedProperties struct {

	// 工作空间ID。
	WorkSpaceId *string `json:"workSpaceId,omitempty" xml:"workSpaceId"`

	// 资源ID。
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId"`

	// 是否是试用集群。
	Trial *string `json:"trial,omitempty" xml:"trial"`
}

func (o ExtendedProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendedProperties struct{}"
	}

	return strings.Join([]string{"ExtendedProperties", string(data)}, " ")
}
