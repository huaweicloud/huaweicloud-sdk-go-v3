package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityGroupTagsRequest Request Object
type ListSecurityGroupTagsRequest struct {
}

func (o ListSecurityGroupTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupTagsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupTagsRequest", string(data)}, " ")
}
