package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Resource struct {

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源详情
	ResourceDetail *interface{} `json:"resource_detail,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// -| 系统标签列表。仅op_service权限才可以可以获取此字段：目前只包含一个resource_tag 结构体key：_sys_enterprise_project_idvalue：企业项目id，0表示默认企业项目非op_service场景不能返回此字段。
	SysTags *[]ResourceTag `json:"sys_tags,omitempty"`
}

func (o Resource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resource struct{}"
	}

	return strings.Join([]string{"Resource", string(data)}, " ")
}
