package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachSubNetworkInterfaceResponse Response Object
type DetachSubNetworkInterfaceResponse struct {

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId *string `json:"request_id,omitempty"`

	SubNetworkInterface *SubNetworkInterface `json:"sub_network_interface,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o DetachSubNetworkInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachSubNetworkInterfaceResponse struct{}"
	}

	return strings.Join([]string{"DetachSubNetworkInterfaceResponse", string(data)}, " ")
}
