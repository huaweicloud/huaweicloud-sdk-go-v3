package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandshakeDto 两个账号（发起者和接收者）之间为了能安全地建立关系，所需要交换的信息。例如，当管理账号（发起者）邀请另一个账号（接收者）加入其组织时，两个账号一系列邀请（握手）请求和响应交换信息。
type HandshakeDto struct {

	// 邀请（握手）的唯一标识符（ID）。源账号在发起邀请（握手）时创建ID。
	Id string `json:"id"`

	// 邀请（握手）的统一资源名称。
	Urn string `json:"urn"`

	// 邀请（握手）请求被接受、取消、拒绝或到期的日期和时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 提出邀请（握手）请求的日期和时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 邀请（握手）过期的日期和时间。
	ExpiredAt *sdktime.SdkTime `json:"expired_at"`

	// 组织管理账号的唯一标识符（ID）。
	ManagementAccountId string `json:"management_account_id"`

	// 组织管理账号的名称。
	ManagementAccountName string `json:"management_account_name"`

	// 组织的唯一标识符（ID）。
	OrganizationId string `json:"organization_id"`

	// 给收件账号所有者的邮件中的附加信息。
	Notes string `json:"notes"`

	Target *TargetDto `json:"target"`

	// 邀请（握手）的当前状态, pending：邀请中；accepted：接受邀请；cancelled：取消邀请；declined：拒绝邀请；expired：邀请过期。
	Status string `json:"status"`
}

func (o HandshakeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandshakeDto struct{}"
	}

	return strings.Join([]string{"HandshakeDto", string(data)}, " ")
}
