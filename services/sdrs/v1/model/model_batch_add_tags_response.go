package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAddTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddTagsResponse", string(data)}, " ")
}
