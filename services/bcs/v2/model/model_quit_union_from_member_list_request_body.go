package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuitUnionFromMemberListRequestBody struct {

	// 邀请方BCS服务实例ID。可调用“查询服务实例列表”接口获取对应的ID
	InviterBcsid string `json:"inviter_bcsid"`

	// 邀请方项目ID。控制台->邀请方帐号->我的凭证->API凭证->项目列表，选择对应的项目ID
	InviterProjectid string `json:"inviter_projectid"`

	// 邀请方租户ID。控制台->邀请方帐号->我的凭证->API凭证->帐号ID
	InviterDomainid string `json:"inviter_domainid"`

	// 邀请方租户名称。控制台->邀请方帐号->我的凭证->API凭证->帐号名
	InviterUsername string `json:"inviter_username"`

	// 联盟通道名称。BCS管理面->成员管理->通道，选择对应的邀请的通道
	ChannelName string `json:"channel_name"`

	// 被邀请方BCS服务实例ID。可调用“查询服务实例列表”接口获取对应的id
	InviteeBcsid string `json:"invitee_bcsid"`

	// 被邀请方项目ID。控制台->被邀请方帐号->我的凭证->API凭证->项目列表，选择对应的项目ID
	InviteeProjectid string `json:"invitee_projectid"`

	// 被邀请方租户ID。控制台->被邀请方帐号->我的凭证->API凭证->帐号ID
	InviteeDomainid string `json:"invitee_domainid"`
}

func (o QuitUnionFromMemberListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuitUnionFromMemberListRequestBody struct{}"
	}

	return strings.Join([]string{"QuitUnionFromMemberListRequestBody", string(data)}, " ")
}
