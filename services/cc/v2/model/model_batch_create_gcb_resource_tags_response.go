package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGcbResourceTagsResponse Response Object
type BatchCreateGcbResourceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateGcbResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGcbResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateGcbResourceTagsResponse", string(data)}, " ")
}
