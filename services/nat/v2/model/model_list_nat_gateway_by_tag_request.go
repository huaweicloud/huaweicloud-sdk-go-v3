package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNatGatewayByTagRequest Request Object
type ListNatGatewayByTagRequest struct {
	Body *ListNatsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListNatGatewayByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewayByTagRequest struct{}"
	}

	return strings.Join([]string{"ListNatGatewayByTagRequest", string(data)}, " ")
}
