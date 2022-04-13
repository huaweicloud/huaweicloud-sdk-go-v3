package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResourceGroupResponse struct {
	// 资源分组的名称，如：Resource-Group-ECS-01。

	GroupName *string `json:"group_name,omitempty"`
	// 资源分组的ID，如：rg1603786526428bWbVmk4rP。

	GroupId *string `json:"group_id,omitempty"`
	// 一组或者多个资源信息。

	Resources *[]ResourceGroup `json:"resources,omitempty"`
	// 资源分组的当前状态，值可为health、unhealth、no_alarm_rule；health表示健康，unhealth表示不健康，no_alarm_rule表示未设置告警规则。

	Status *string `json:"status,omitempty"`
	// 资源分组的创建时间，UNIX时间戳，单位毫秒；如：1603819753000。

	CreateTime *int64 `json:"create_time,omitempty"`

	MetaData *MetaData `json:"meta_data,omitempty"`
	// 创建资源分组时关联的企业项目，默认值为0，表示企业项目为default。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceGroupResponse", string(data)}, " ")
}
