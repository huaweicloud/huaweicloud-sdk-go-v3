package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteMemberInviteRequestBody struct {
	// 邀请实例id

	BcsId string `json:"bcs_id"`
	// 邀请加入的通道名

	ChannelName string `json:"channel_name"`
	// 被邀请的用户列表，对应信息可通过获取联盟成员列表（ListMembers）接口查询，或被邀请方已加入联盟，或邀请状态为released时，需填写准确的被邀请方bcs实例id和邀请状态

	InvitedUserinfo []InvitationDetail `json:"invited_userinfo"`
}

func (o DeleteMemberInviteRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberInviteRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteMemberInviteRequestBody", string(data)}, " ")
}
