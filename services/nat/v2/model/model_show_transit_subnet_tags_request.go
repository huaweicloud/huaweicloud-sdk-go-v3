package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTransitSubnetTagsRequest Request Object
type ShowTransitSubnetTagsRequest struct {

	// 中转子网的ID。
	ResourceId string `json:"resource_id"`
}

func (o ShowTransitSubnetTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTransitSubnetTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowTransitSubnetTagsRequest", string(data)}, " ")
}
