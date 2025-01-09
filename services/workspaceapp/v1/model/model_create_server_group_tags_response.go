package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServerGroupTagsResponse Response Object
type CreateServerGroupTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateServerGroupTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServerGroupTagsResponse struct{}"
	}

	return strings.Join([]string{"CreateServerGroupTagsResponse", string(data)}, " ")
}
