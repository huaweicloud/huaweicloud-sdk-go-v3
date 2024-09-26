package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNetworkInstanceResponse Response Object
type ShowNetworkInstanceResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	NetworkInstance *NetworkInstance `json:"network_instance"`
	HttpStatusCode  int              `json:"-"`
}

func (o ShowNetworkInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNetworkInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowNetworkInstanceResponse", string(data)}, " ")
}
