package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCentralNetworkResponse Response Object
type DeleteCentralNetworkResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetwork *CentralNetwork `json:"central_network"`
	HttpStatusCode int             `json:"-"`
}

func (o DeleteCentralNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCentralNetworkResponse struct{}"
	}

	return strings.Join([]string{"DeleteCentralNetworkResponse", string(data)}, " ")
}
