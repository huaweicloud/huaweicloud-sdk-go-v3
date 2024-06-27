package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SimpleResourceInfo 反选的资源列表
type SimpleResourceInfo struct {

	// 资源类型
	Type *string `json:"type,omitempty"`

	// 资源责任人
	Owner *string `json:"owner,omitempty"`

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o SimpleResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleResourceInfo struct{}"
	}

	return strings.Join([]string{"SimpleResourceInfo", string(data)}, " ")
}
