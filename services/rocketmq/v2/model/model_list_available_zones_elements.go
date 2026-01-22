package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAvailableZonesElements struct {

	// **参数解释**： 是否售罄。 **约束限制**： 不涉及。 **取值范围**： - true：售罄。 - false：未售罄。 **默认取值**： 不涉及。
	SoldOut *bool `json:"soldOut,omitempty"`

	// **参数解释**： 可用区ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 可用区编码。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**： 可用区名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 可用区端口号。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Port *string `json:"port,omitempty"`

	// **参数解释**： 分区上是否还有可用资源。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ResourceAvailability *string `json:"resource_availability,omitempty"`

	// **参数解释**： 是否为默认可用区。 **约束限制**： 不涉及。 **取值范围**： - true：默认可用区 - false：不是默认可用区 **默认取值**： 不涉及。
	DefaultAz *bool `json:"default_az,omitempty"`

	// **参数解释**： 剩余时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RemainTime *int64 `json:"remain_time,omitempty"`

	// **参数解释**： 是否支持IPv6。[华为云Stack不支持此参数。](tag:hcs,hcs_oemout) **约束限制**： 不涉及。 **取值范围**： - true：支持 - false：不支持 **默认取值**： 不涉及。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// **参数解释**： AZ模式。 **取值范围**： - dedicated：边缘AZ - shared：本地AZ - center：中心AZ
	Mode *string `json:"mode,omitempty"`

	// **参数解释**： 可用区分组。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Scope *string `json:"scope,omitempty"`
}

func (o ListAvailableZonesElements) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesElements struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesElements", string(data)}, " ")
}
