package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagCentralNetworkRequest Request Object
type TagCentralNetworkRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`

	Body *TagCentralNetworkRequestBody `json:"body,omitempty"`
}

func (o TagCentralNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagCentralNetworkRequest struct{}"
	}

	return strings.Join([]string{"TagCentralNetworkRequest", string(data)}, " ")
}
