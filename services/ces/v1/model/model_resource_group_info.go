package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceGroupInfo 资源分组信息
type ResourceGroupInfo struct {

	// 资源分组的名称，如：ResourceGroup-Test01。
	GroupName *string `json:"group_name,omitempty"`

	// 资源添加/匹配方式，取值只能为EPS（匹配企业项目）,TAG（匹配标签）,NAME（匹配实例名称）, COMB（组合匹配）,Manual/空值（手动添加）
	Type *string `json:"type,omitempty"`

	// 企业项目ID列表
	RelationIds *[]string `json:"relation_ids,omitempty"`

	// 资源分组的ID，如：rg1603786526428bWbVmk4rP。
	GroupId *string `json:"group_id,omitempty"`

	// 资源分组的创建时间，UNIX时间戳，单位毫秒；如：1603819753000。
	CreateTime *int64 `json:"create_time,omitempty"`

	InstanceStatistics *InstanceStatistics `json:"instance_statistics,omitempty"`

	Status *StatusSchema `json:"status,omitempty"`

	// 创建资源分组时关联的企业项目，默认值为0，表示企业项目为default。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 一组或者多个资源信息，默认为空。
	Resources *[]Resource `json:"resources,omitempty"`
}

func (o ResourceGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroupInfo struct{}"
	}

	return strings.Join([]string{"ResourceGroupInfo", string(data)}, " ")
}
