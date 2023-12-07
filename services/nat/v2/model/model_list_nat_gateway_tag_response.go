package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNatGatewayTagResponse Response Object
type ListNatGatewayTagResponse struct {

	// 标签列表。
	Tags           *[]TagsBody `json:"tags,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListNatGatewayTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewayTagResponse struct{}"
	}

	return strings.Join([]string{"ListNatGatewayTagResponse", string(data)}, " ")
}
