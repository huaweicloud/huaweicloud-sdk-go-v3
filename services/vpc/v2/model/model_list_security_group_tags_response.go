package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityGroupTagsResponse Response Object
type ListSecurityGroupTagsResponse struct {

	// tag对象列表
	Tags           *[]ListTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListSecurityGroupTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupTagsResponse", string(data)}, " ")
}
