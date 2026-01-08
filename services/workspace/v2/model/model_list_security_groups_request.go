package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityGroupsRequest Request Object
type ListSecurityGroupsRequest struct {
}

func (o ListSecurityGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupsRequest", string(data)}, " ")
}
