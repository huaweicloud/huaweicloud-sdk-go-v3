package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkRequestBody 更新中心网络的请求体。
type UpdateCentralNetworkRequestBody struct {
	CentralNetwork *UpdateCentralNetwork `json:"central_network"`
}

func (o UpdateCentralNetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkRequestBody", string(data)}, " ")
}
