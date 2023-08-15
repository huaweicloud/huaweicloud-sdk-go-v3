package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateResourceTagsResponse Response Object
type BatchCreateResourceTagsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateResourceTagsResponse", string(data)}, " ")
}
