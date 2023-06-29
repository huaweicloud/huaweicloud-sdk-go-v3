package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateListenerTagsResponse Response Object
type CreateListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateListenerTagsResponse", string(data)}, " ")
}
