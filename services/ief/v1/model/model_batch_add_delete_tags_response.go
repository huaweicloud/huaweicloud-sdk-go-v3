package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddDeleteTagsResponse Response Object
type BatchAddDeleteTagsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddDeleteTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDeleteTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddDeleteTagsResponse", string(data)}, " ")
}
