package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerOfferingsResponse Response Object
type ListServerOfferingsResponse struct {

	// 商品列表
	Offerings *[]ServerOffering `json:"offerings,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListServerOfferingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerOfferingsResponse struct{}"
	}

	return strings.Join([]string{"ListServerOfferingsResponse", string(data)}, " ")
}
