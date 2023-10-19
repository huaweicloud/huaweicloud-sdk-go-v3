package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkRequest Request Object
type UpdateCentralNetworkRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	Body *UpdateCentralNetworkRequestBody `json:"body,omitempty"`
}

func (o UpdateCentralNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkRequest struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkRequest", string(data)}, " ")
}
