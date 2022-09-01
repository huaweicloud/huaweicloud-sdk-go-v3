package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 邀请联盟成员
type BatchInviteMembersToChannelRequestBody struct {

	// 发出邀请的BCS服务实例id
	BcsId string `json:"bcs_id" xml:"bcs_id"`

	// 邀请加入的通道名
	ChannelName string `json:"channel_name" xml:"channel_name"`

	// 发出邀请的租户名
	InvitorUsername *string `json:"invitor_username,omitempty" xml:"invitor_username"`

	// 被邀请的用户列表
	InvitedUserinfo []InvitedDomain `json:"invited_userinfo" xml:"invited_userinfo"`
}

func (o BatchInviteMembersToChannelRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInviteMembersToChannelRequestBody struct{}"
	}

	return strings.Join([]string{"BatchInviteMembersToChannelRequestBody", string(data)}, " ")
}
