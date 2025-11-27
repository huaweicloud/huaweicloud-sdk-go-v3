package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDeleteTransitSubnetTagsResponse Response Object
type BatchCreateDeleteTransitSubnetTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateDeleteTransitSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteTransitSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteTransitSubnetTagsResponse", string(data)}, " ")
}
