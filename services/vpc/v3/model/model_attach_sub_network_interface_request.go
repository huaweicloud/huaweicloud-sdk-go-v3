package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachSubNetworkInterfaceRequest Request Object
type AttachSubNetworkInterfaceRequest struct {

	// **参数解释**： 辅助弹性网卡的资源ID。 **取值范围**： 不涉及。
	SubNetworkInterfaceId string `json:"sub_network_interface_id"`

	Body *AttachSubNetworkInterfaceRequestBody `json:"body,omitempty"`
}

func (o AttachSubNetworkInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachSubNetworkInterfaceRequest struct{}"
	}

	return strings.Join([]string{"AttachSubNetworkInterfaceRequest", string(data)}, " ")
}
