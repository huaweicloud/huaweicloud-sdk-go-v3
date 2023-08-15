package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDeleteTransitIpTagsResponse Response Object
type BatchCreateDeleteTransitIpTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateDeleteTransitIpTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteTransitIpTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteTransitIpTagsResponse", string(data)}, " ")
}
