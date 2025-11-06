package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceGroupResponse Response Object
type ShowResourceGroupResponse struct {

	// **参数解释** 资源分组的名称 **约束限制** 不涉及 **取值范围** 只能为字母、数字、汉字、-或_，长度为[1,128]个字符 **默认取值** 不涉及
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 资源分组ID     **约束限制**： 不涉及。  **取值范围**： 以rg开头，后跟22位由字母或数字组成的字符串。长度为[2,24]个字符。       **默认取值**： 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释** 创建的资源分组选择一个或者多个资源。 **约束限制** 不超过1000个资源。
	Resources *[]ResourceGroup `json:"resources,omitempty"`

	Status *StatusSchema `json:"status,omitempty"`

	// **参数解释**： 资源分组的创建时间，UNIX时间戳，单位毫秒；如：1603819753000。     **约束限制**： 不涉及。  **取值范围**： 在[1,9223372036854775807]区间内 **默认取值**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	MetaData *MetaData `json:"meta_data,omitempty"`

	// **参数解释**： 创建资源分组时关联的企业项目，默认值为0，表示企业项目为default。     **约束限制**： 不涉及。  **取值范围**： 由数字、字母和-组成，或者为0（默认企业项目ID）。 **默认取值**： 不涉及。
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
