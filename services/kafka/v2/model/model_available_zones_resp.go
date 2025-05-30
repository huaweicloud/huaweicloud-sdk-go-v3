package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableZonesResp struct {

	// **参数解释**： 是否售罄。 **取值范围**： - true：售罄 - false：没有售罄
	SoldOut *bool `json:"soldOut,omitempty"`

	// **参数解释**： 可用区ID。 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 可用区编码。 **取值范围**： 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**： 可用区名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 可用区端口号。 **取值范围**： 不涉及
	Port *string `json:"port,omitempty"`

	// **参数解释**： 可用区上是否还有可用资源。 **取值范围**： - true：有可用资源 - false：无可用资源
	ResourceAvailability *string `json:"resource_availability,omitempty"`

	// **参数解释**： 是否为默认可用区。 **取值范围**： - true：默认可用区 - false：不是默认可用区
	DefaultAz *bool `json:"default_az,omitempty"`

	// **参数解释**： 剩余时间，以Unix时间戳显示。 **取值范围**： 不涉及
	RemainTime *int64 `json:"remain_time,omitempty"`

	// **参数解释**： 是否支持IPv6。 **取值范围**： - true：支持 - false：不支持
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`
}

func (o AvailableZonesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableZonesResp struct{}"
	}

	return strings.Join([]string{"AvailableZonesResp", string(data)}, " ")
}
