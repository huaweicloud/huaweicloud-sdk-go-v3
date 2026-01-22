package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateNetworkSegmentsResponse Response Object
type ListPrivateNetworkSegmentsResponse struct {
	Data           *ListPrivateNetworkSegmentsVo `json:"data,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListPrivateNetworkSegmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNetworkSegmentsResponse struct{}"
	}

	return strings.Join([]string{"ListPrivateNetworkSegmentsResponse", string(data)}, " ")
}
