package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateServerTagsResponse Response Object
type BatchCreateServerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateServerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateServerTagsResponse", string(data)}, " ")
}
