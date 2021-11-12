package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchCreateListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateListenerTagsResponse", string(data)}, " ")
}
