package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IsMemberInGroupsResponse Response Object
type IsMemberInGroupsResponse struct {

	// 成员是否在用户组内的结果列表
	Results        *[]GroupMembershipExistenceResultDto `json:"results,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o IsMemberInGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsMemberInGroupsResponse struct{}"
	}

	return strings.Join([]string{"IsMemberInGroupsResponse", string(data)}, " ")
}
