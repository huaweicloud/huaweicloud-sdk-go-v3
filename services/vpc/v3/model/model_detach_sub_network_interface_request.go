package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachSubNetworkInterfaceRequest Request Object
type DetachSubNetworkInterfaceRequest struct {

	// **参数解释**： 辅助弹性网卡的资源ID。 **取值范围**： 不涉及。
	SubNetworkInterfaceId string `json:"sub_network_interface_id"`
}

func (o DetachSubNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachSubNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"DetachSubNetworkInterfaceRequest", string(data)}, " ")
}
