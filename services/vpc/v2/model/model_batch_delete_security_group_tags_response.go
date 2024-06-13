package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityGroupTagsResponse Response Object
type BatchDeleteSecurityGroupTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSecurityGroupTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityGroupTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityGroupTagsResponse", string(data)}, " ")
}
