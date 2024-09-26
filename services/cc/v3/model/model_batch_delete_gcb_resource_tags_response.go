package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteGcbResourceTagsResponse Response Object
type BatchDeleteGcbResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteGcbResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteGcbResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteGcbResourceTagsResponse", string(data)}, " ")
}
