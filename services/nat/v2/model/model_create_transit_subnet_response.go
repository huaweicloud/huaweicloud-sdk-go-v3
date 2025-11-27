package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransitSubnetResponse Response Object
type CreateTransitSubnetResponse struct {
	TransitSubnet *TransitSubnet `json:"transit_subnet,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTransitSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitSubnetResponse struct{}"
	}

	return strings.Join([]string{"CreateTransitSubnetResponse", string(data)}, " ")
}
