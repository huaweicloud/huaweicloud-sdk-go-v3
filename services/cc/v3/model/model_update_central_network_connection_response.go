package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkConnectionResponse Response Object
type UpdateCentralNetworkConnectionResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetworkConnection *CentralNetworkConnection `json:"central_network_connection"`
	HttpStatusCode           int                       `json:"-"`
}

func (o UpdateCentralNetworkConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkConnectionResponse", string(data)}, " ")
}
