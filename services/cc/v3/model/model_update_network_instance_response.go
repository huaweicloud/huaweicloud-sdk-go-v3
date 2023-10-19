package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNetworkInstanceResponse Response Object
type UpdateNetworkInstanceResponse struct {

	// 资源ID标识符。
	RequestId string `json:"request_id"`

	NetworkInstance *NetworkInstance `json:"network_instance"`
	HttpStatusCode  int              `json:"-"`
}

func (o UpdateNetworkInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNetworkInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateNetworkInstanceResponse", string(data)}, " ")
}
