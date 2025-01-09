package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizationMail 邮件记录详细信息。
type AuthorizationMail struct {

	// 邮件记录id。
	Id string `json:"id"`

	// 用户(组)。
	Account *string `json:"account,omitempty"`

	AccountAuthType *AccountTypeEnum `json:"account_auth_type,omitempty"`

	// 授权对象名称。
	AccountAuthName *string `json:"account_auth_name,omitempty"`

	// 应用组ID。
	AppGroupId *string `json:"app_group_id,omitempty"`

	// 应用组名称。
	AppGroupName *string `json:"app_group_name,omitempty"`

	// 授权类型： - ADD_GROUP_AUTHORIZATION 添加组授权邮件 - DEL_GROUP_AUTHORIZATION 删除组授权邮件 - ADD_GROUP_AUTHORIZATION_SMS 添加组授权短信 - DEL_GROUP_AUTHORIZATION_SMS 删除组授权短信
	MailSendType *string `json:"mail_send_type,omitempty"`

	// 发送结果。 - SUEECESS -发送成功 - FAIL -发送失败
	MailSendResult *string `json:"mail_send_result,omitempty"`

	// 报错信息。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 发布时间。
	SendAt *sdktime.SdkTime `json:"send_at,omitempty"`
}

func (o AuthorizationMail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizationMail struct{}"
	}

	return strings.Join([]string{"AuthorizationMail", string(data)}, " ")
}
