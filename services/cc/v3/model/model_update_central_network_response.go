package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkResponse Response Object
type UpdateCentralNetworkResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetwork *CentralNetwork `json:"central_network"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateCentralNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkResponse struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkResponse", string(data)}, " ")
}
