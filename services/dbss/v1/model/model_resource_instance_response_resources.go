package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceInstanceResponseResources struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源名称，资源没有名称时默认为空字符串，eip返回ip地址。
	ResourceName string `json:"resource_name"`

	// 资源详情。 资源对象，用于扩展，默认为空。
	ResourceDetail *interface{} `json:"resource_detail"`

	// 标签列表，没有标签默认为空数组
	Tags []ResourceInstanceResponseTags `json:"tags"`

	// 仅op_service权限才可以可以获取此字段： 目前只包含一个resource_tag 结构体 key：_sys_enterprise_project_id value：企业项目id，0表示默认企业项目 非op_service场景不能返回此字段
	SysTags *[]ResourceInstanceResponseSysTags `json:"sys_tags,omitempty"`
}

func (o ResourceInstanceResponseResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstanceResponseResources struct{}"
	}

	return strings.Join([]string{"ResourceInstanceResponseResources", string(data)}, " ")
}
