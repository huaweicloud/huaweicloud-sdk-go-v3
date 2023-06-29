package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerResource
type ServerResource struct {

	// 云服务器ID。
	ResourceId string `json:"resource_id"`

	// 预留字段。
	ResourceDetail string `json:"resource_detail"`

	// 标签列表
	Tags []ResourceTag `json:"tags"`

	// 资源名称，即弹性云服务器名称。
	ResourceName string `json:"resource_name"`
}

func (o ServerResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerResource struct{}"
	}

	return strings.Join([]string{"ServerResource", string(data)}, " ")
}
