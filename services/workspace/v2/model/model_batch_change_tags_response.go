package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeTagsResponse Response Object
type BatchChangeTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchChangeTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchChangeTagsResponse", string(data)}, " ")
}
