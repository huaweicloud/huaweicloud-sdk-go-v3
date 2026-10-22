package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteV2ResourceTagsResponse Response Object
type BatchDeleteV2ResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteV2ResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteV2ResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteV2ResourceTagsResponse", string(data)}, " ")
}
