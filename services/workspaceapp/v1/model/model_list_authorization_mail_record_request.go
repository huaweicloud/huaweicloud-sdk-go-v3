package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuthorizationMailRecordRequest Request Object
type ListAuthorizationMailRecordRequest struct {

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`

	// 用户(组)名称。
	Account *string `json:"account,omitempty"`

	// 授权类型： - ADD_GROUP_AUTHORIZATION 添加组授权 - DEL_GROUP_AUTHORIZATION 删除组授权
	MailSendType *string `json:"mail_send_type,omitempty"`

	// 邮件发送结果(SUCCESS|FAIL)。
	MailSendResult *string `json:"mail_send_result,omitempty"`
}

func (o ListAuthorizationMailRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizationMailRecordRequest struct{}"
	}

	return strings.Join([]string{"ListAuthorizationMailRecordRequest", string(data)}, " ")
}
