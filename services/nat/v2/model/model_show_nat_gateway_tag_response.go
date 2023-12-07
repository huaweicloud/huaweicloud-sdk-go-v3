package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNatGatewayTagResponse Response Object
type ShowNatGatewayTagResponse struct {

	// 标签列表。
	Tags           *[]TagBody `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowNatGatewayTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNatGatewayTagResponse struct{}"
	}

	return strings.Join([]string{"ShowNatGatewayTagResponse", string(data)}, " ")
}
