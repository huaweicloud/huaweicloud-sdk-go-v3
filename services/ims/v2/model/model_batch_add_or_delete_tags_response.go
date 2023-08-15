package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddOrDeleteTagsResponse Response Object
type BatchAddOrDeleteTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchAddOrDeleteTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrDeleteTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddOrDeleteTagsResponse", string(data)}, " ")
}
