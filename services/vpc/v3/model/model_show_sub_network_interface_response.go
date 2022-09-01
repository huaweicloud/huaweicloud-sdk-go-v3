package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSubNetworkInterfaceResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	SubNetworkInterface *SubNetworkInterface `json:"sub_network_interface,omitempty" xml:"sub_network_interface"`
	HttpStatusCode      int                  `json:"-"`
}

func (o ShowSubNetworkInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubNetworkInterfaceResponse struct{}"
	}

	return strings.Join([]string{"ShowSubNetworkInterfaceResponse", string(data)}, " ")
}
