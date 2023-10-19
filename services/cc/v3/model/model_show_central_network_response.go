package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCentralNetworkResponse Response Object
type ShowCentralNetworkResponse struct {

	// 资源ID标识符。
	RequestId string `json:"request_id"`

	CentralNetwork *CentralNetwork `json:"central_network"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowCentralNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCentralNetworkResponse struct{}"
	}

	return strings.Join([]string{"ShowCentralNetworkResponse", string(data)}, " ")
}
