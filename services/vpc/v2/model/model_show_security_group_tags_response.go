package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityGroupTagsResponse Response Object
type ShowSecurityGroupTagsResponse struct {

	// tag对象列表
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowSecurityGroupTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityGroupTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityGroupTagsResponse", string(data)}, " ")
}
