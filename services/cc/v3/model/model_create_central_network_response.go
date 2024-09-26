package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkResponse Response Object
type CreateCentralNetworkResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	CentralNetwork *CentralNetwork `json:"central_network"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateCentralNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkResponse struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkResponse", string(data)}, " ")
}
