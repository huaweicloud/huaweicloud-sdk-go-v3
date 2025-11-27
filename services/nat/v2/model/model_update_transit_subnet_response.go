package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTransitSubnetResponse Response Object
type UpdateTransitSubnetResponse struct {
	TransitSubnet *TransitSubnet `json:"transit_subnet,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTransitSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransitSubnetResponse struct{}"
	}

	return strings.Join([]string{"UpdateTransitSubnetResponse", string(data)}, " ")
}
