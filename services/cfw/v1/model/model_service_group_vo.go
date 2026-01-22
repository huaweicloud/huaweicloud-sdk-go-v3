package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceGroupVo struct {

	// **参数解释**： 服务（协议、源端口、目的端口）组的名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 协议列表 **约束限制**： 不涉及 **取值范围**： - 6：TCP - 17：UDP - 1：ICMP - 58：ICMPV6 - -1：Any **默认取值**： 不涉及
	Protocols *[]int32 `json:"protocols,omitempty"`

	// **参数解释**： 服务（协议、源端口、目的端口）组的类型 **约束限制**： 不涉及 **取值范围**： 0表示自定义服务组，1表示预定义服务组 **默认取值**： 不涉及
	ServiceSetType *int32 `json:"service_set_type,omitempty"`

	// **参数解释**： 服务组ID，可通过[获取服务组列表接口](ListServiceSets.xml)查询获得，通过返回值中的data.records.set_id（.表示各对象之间层级的区分）获得。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	SetId *string `json:"set_id,omitempty"`
}

func (o ServiceGroupVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceGroupVo struct{}"
	}

	return strings.Join([]string{"ServiceGroupVo", string(data)}, " ")
}
