package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteListenerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteListenerTagsResponse", string(data)}, " ")
}
