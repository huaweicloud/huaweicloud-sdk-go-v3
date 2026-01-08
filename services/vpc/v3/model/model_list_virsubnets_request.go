package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVirsubnetsRequest Request Object
type ListVirsubnetsRequest struct {

	// **参数解释**： 每页返回的资源个数。 **取值范围**： 0~2000。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页查询起始的资源ID，为空时查询第一页。 **取值范围**： 子网的资源ID。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**： 子网的资源ID，填写后按照ID进行过滤，支持传入多个ID过滤。 **取值范围**： 不涉及。
	Id *[]string `json:"id,omitempty"`

	// **参数解释**： 子网的名称，填写后按照名称进行过滤，支持传入多个名称过滤。 **取值范围**： 不涉及。
	Name *[]string `json:"name,omitempty"`

	// **参数解释**： 子网所在VPC的资源ID，填写后按照VPC资源ID进行过滤，支持传入多个ID过滤。 **取值范围**： 不涉及。
	VpcId *[]string `json:"vpc_id,omitempty"`

	// **参数解释**： 子网的状态，填写后按照状态进行过滤，只支持传入单个状态过滤。 **取值范围**： - ACTIVE：表示子网已挂载到VPC上。 - UNKNOWN：表示子网还未挂载到VPC上。 - ERROR：表示子网状态故障。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 子网的作用域，填写后按照作用域进行过滤，支持传入多个作用域过滤。 **取值范围**： - center：表示作用域为中心。 - {publicBorderGroup}：表示作用域为具体的公网边界组。公网边界组限制子网的可用区范围，可关联多个边缘可用区。
	Scope *[]string `json:"scope,omitempty"`

	// **参数解释**： 子网的可用区ID，填写后按照可用区ID进行过滤，支持传入多个可用区ID过滤。 **取值范围**： 不涉及
	ZoneId *[]string `json:"zone_id,omitempty"`

	// **参数解释**： 子网的描述信息，填写后按照描述信息进行过滤，支持传入多个描述信息过滤。 **取值范围**： 不涉及。
	Description *[]string `json:"description,omitempty"`
}

func (o ListVirsubnetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirsubnetsRequest struct{}"
	}

	return strings.Join([]string{"ListVirsubnetsRequest", string(data)}, " ")
}
