package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerGroupTagsResponse Response Object
type DeleteServerGroupTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerGroupTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupTagsResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupTagsResponse", string(data)}, " ")
}
