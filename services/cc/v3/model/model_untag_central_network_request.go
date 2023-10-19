package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UntagCentralNetworkRequest Request Object
type UntagCentralNetworkRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	Body *UntagCentralNetworkRequestBody `json:"body,omitempty"`
}

func (o UntagCentralNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagCentralNetworkRequest struct{}"
	}

	return strings.Join([]string{"UntagCentralNetworkRequest", string(data)}, " ")
}
