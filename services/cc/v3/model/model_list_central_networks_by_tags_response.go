package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCentralNetworksByTagsResponse Response Object
type ListCentralNetworksByTagsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 中心网络列表。
	CentralNetworks []CentralNetwork `json:"central_networks"`
	HttpStatusCode  int              `json:"-"`
}

func (o ListCentralNetworksByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCentralNetworksByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListCentralNetworksByTagsResponse", string(data)}, " ")
}
