package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceShareInvitation 资源共享邀请的详细信息。
type ResourceShareInvitation struct {

	// 接收资源共享邀请的账号ID。
	ReceiverAccountId *string `json:"receiver_account_id,omitempty"`

	// 资源共享实例的ID。
	ResourceShareId *string `json:"resource_share_id,omitempty"`

	// 资源共享实例的名称。
	ResourceShareName *string `json:"resource_share_name,omitempty"`

	// 资源共享邀请的ID。
	ResourceShareInvitationId *string `json:"resource_share_invitation_id,omitempty"`

	// 发送资源共享邀请的账号ID。
	SenderAccountId *string `json:"sender_account_id,omitempty"`

	// 资源共享邀请的当前状态。
	Status *string `json:"status,omitempty"`

	// 创建邀请的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 最后一次更新邀请的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ResourceShareInvitation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceShareInvitation struct{}"
	}

	return strings.Join([]string{"ResourceShareInvitation", string(data)}, " ")
}
