package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachSubNetworkInterfaceResponse Response Object
type AttachSubNetworkInterfaceResponse struct {

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId *string `json:"request_id,omitempty"`

	SubNetworkInterface *SubNetworkInterface `json:"sub_network_interface,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o AttachSubNetworkInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachSubNetworkInterfaceResponse struct{}"
	}

	return strings.Join([]string{"AttachSubNetworkInterfaceResponse", string(data)}, " ")
}
