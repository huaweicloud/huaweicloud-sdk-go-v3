package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServerGroupTagsResponse Response Object
type BatchDeleteServerGroupTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteServerGroupTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerGroupTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerGroupTagsResponse", string(data)}, " ")
}
