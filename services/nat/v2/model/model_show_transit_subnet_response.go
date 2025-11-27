package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransitSubnetResponse Response Object
type ShowTransitSubnetResponse struct {
	TransitSubnet *TransitSubnet `json:"transit_subnet,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTransitSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitSubnetResponse struct{}"
	}

	return strings.Join([]string{"ShowTransitSubnetResponse", string(data)}, " ")
}
