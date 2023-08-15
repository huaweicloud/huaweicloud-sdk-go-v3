package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSubnetTagsResponse Response Object
type BatchCreateSubnetTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSubnetTagsResponse", string(data)}, " ")
}
