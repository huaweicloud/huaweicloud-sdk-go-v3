package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateNatTagsRequest Request Object
type ShowPrivateNatTagsRequest struct {

	// 私网NAT网关的ID。
	ResourceId string `json:"resource_id"`
}

func (o ShowPrivateNatTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateNatTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateNatTagsRequest", string(data)}, " ")
}
