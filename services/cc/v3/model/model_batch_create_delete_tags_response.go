package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDeleteTagsResponse Response Object
type BatchCreateDeleteTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateDeleteTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDeleteTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDeleteTagsResponse", string(data)}, " ")
}
