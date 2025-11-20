package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MasterEipResponseStatus **参数解释**： 状态信息 **约束限制**： 不涉及
type MasterEipResponseStatus struct {

	// **参数解释**： 集群访问的PrivateIP(HA集群返回VIP) **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	PrivateEndpoint *string `json:"privateEndpoint,omitempty"`

	// **参数解释**： 集群访问的PublicIP **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	PublicEndpoint *string `json:"publicEndpoint,omitempty"`
}

func (o MasterEipResponseStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipResponseStatus struct{}"
	}

	return strings.Join([]string{"MasterEipResponseStatus", string(data)}, " ")
}
