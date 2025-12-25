package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceInfo struct {

	// 资产ID
	ResourceId string `json:"resource_id"`

	// 标签
	Tags *[]Match `json:"tags,omitempty"`

	// 系统标签
	SysTags *[]Match `json:"sys_tags,omitempty"`

	// 资产名称
	ResourceName string `json:"resource_name"`

	// 资产细节
	ResourceDetail *interface{} `json:"resource_detail,omitempty"`
}

func (o ResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInfo struct{}"
	}

	return strings.Join([]string{"ResourceInfo", string(data)}, " ")
}
