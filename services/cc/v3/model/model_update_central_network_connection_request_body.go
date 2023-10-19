package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkConnectionRequestBody 更新中心网络连接信息的请求体。
type UpdateCentralNetworkConnectionRequestBody struct {
	CentralNetworkConnection *UpdateCentralNetworkConnection `json:"central_network_connection"`
}

func (o UpdateCentralNetworkConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkConnectionRequestBody", string(data)}, " ")
}
