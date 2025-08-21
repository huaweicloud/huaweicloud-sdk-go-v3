package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationTreeRequest Request Object
type ListOrganizationTreeRequest struct {

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`

	// **参数解释**： 父节点（根或组织单元）的唯一标识符（ID） **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**： 查询返回记录的数量限制 **约束限制**： 不涉及 **取值范围**： 1-2000 **默认取值**： 2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页标记 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Marker *string `json:"marker,omitempty"`
}

func (o ListOrganizationTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationTreeRequest struct{}"
	}

	return strings.Join([]string{"ListOrganizationTreeRequest", string(data)}, " ")
}
