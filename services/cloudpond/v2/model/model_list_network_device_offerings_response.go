package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworkDeviceOfferingsResponse Response Object
type ListNetworkDeviceOfferingsResponse struct {

	// 商品列表
	Offerings *[]NetworkDeviceOffering `json:"offerings,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListNetworkDeviceOfferingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkDeviceOfferingsResponse struct{}"
	}

	return strings.Join([]string{"ListNetworkDeviceOfferingsResponse", string(data)}, " ")
}
