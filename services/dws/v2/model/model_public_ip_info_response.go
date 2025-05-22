package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicIpInfoResponse **参数解释**： 公网IP信息详情。 **取值范围**： 不涉及。
type PublicIpInfoResponse struct {

	// **参数解释**： 公网IP的ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 公网IP。 **取值范围**： 合法的公网IPV4地址。
	Address *string `json:"address,omitempty"`

	// **参数解释**： 公网IP状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 绑定的DWS集群的节点ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 绑定的DWS集群的节点名称。 **取值范围**： 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**： 公网IP带宽信息。 **取值范围**： 不涉及。
	BandwidthSize *string `json:"bandwidth_size,omitempty"`
}

func (o PublicIpInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicIpInfoResponse struct{}"
	}

	return strings.Join([]string{"PublicIpInfoResponse", string(data)}, " ")
}
