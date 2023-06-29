package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreatePublicipTagsResponse Response Object
type BatchCreatePublicipTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreatePublicipTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicipTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicipTagsResponse", string(data)}, " ")
}
