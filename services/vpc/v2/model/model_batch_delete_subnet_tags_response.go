package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSubnetTagsResponse Response Object
type BatchDeleteSubnetTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSubnetTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubnetTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubnetTagsResponse", string(data)}, " ")
}
