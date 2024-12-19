package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDirectConnectLocationsResponse Response Object
type ListDirectConnectLocationsResponse struct {

	// 专线接入点列表
	DirectConnectLocations *[]DirectConnectLocationEntry `json:"direct_connect_locations,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDirectConnectLocationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectConnectLocationsResponse struct{}"
	}

	return strings.Join([]string{"ListDirectConnectLocationsResponse", string(data)}, " ")
}
