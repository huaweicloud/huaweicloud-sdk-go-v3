package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateSecurityGroupTagsResponse Response Object
type BatchCreateSecurityGroupTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchCreateSecurityGroupTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSecurityGroupTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateSecurityGroupTagsResponse", string(data)}, " ")
}
