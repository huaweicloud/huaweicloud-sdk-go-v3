package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteVpcTagsResponse Response Object
type BatchDeleteVpcTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteVpcTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVpcTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteVpcTagsResponse", string(data)}, " ")
}
