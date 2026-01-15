package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityGroupsResponse Response Object
type ListSecurityGroupsResponse struct {

	// 安全组列表。
	SecurityGroups *[]SimpleSecurityGroupsInfo `json:"security_groups,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListSecurityGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupsResponse", string(data)}, " ")
}
