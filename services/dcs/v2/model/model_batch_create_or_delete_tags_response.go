package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateOrDeleteTagsResponse Response Object
type BatchCreateOrDeleteTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateOrDeleteTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteTagsResponse", string(data)}, " ")
}
