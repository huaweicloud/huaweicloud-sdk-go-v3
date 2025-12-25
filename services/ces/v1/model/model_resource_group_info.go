package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceGroupInfo **参数解释** 资源分组信息。
type ResourceGroupInfo struct {

	// **参数解释** 资源分组名称。 **取值范围** 包含字母、数字、_、-或汉字，长度为[1,128]个字符。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释** 资源添加/匹配方式。 **取值范围** 取值只能为EPS（匹配企业项目），TAG（匹配标签），NAME（匹配实例名称），COMB（组合匹配），Manual/空值（手动添加）。
	Type *string `json:"type,omitempty"`

	// **参数解释** 企业项目ID列表。
	RelationIds *[]string `json:"relation_ids,omitempty"`

	// **参数解释**： 资源分组ID。  **取值范围**： 以rg开头，后跟22位由字母或数字组成的字符串。长度为[2,24]个字符。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释** 资源分组的创建时间，UNIX时间戳，单位毫秒；如：1603819753000。 **取值范围** 在[0,9223372036854775807]区间内
	CreateTime *int64 `json:"create_time,omitempty"`

	InstanceStatistics *InstanceStatistics `json:"instance_statistics,omitempty"`

	Status *StatusSchemaResp `json:"status,omitempty"`

	// **参数解释** 资源分组归属企业项目ID。 **取值范围** 由数字、字母和-组成，或者为0（默认企业项目ID）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 一组或者多个资源信息，默认为空。
	Resources *[]Resource `json:"resources,omitempty"`
}

func (o ResourceGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroupInfo struct{}"
	}

	return strings.Join([]string{"ResourceGroupInfo", string(data)}, " ")
}
