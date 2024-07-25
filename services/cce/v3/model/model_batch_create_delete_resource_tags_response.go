package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDeleteResourceTagsResponse Response Object
type BatchCreateDeleteResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateDeleteResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteResourceTagsResponse", string(data)}, " ")
}
