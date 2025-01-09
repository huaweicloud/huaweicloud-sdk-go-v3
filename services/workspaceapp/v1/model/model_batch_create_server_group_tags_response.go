package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateServerGroupTagsResponse Response Object
type BatchCreateServerGroupTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateServerGroupTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateServerGroupTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateServerGroupTagsResponse", string(data)}, " ")
}
