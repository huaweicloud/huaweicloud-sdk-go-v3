package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 变更操作人
type VersionHistoryOperator struct {

	// 用户uuid
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`

	// 租户uuid
	DomainId *string `json:"domain_id,omitempty"`

	// 租户名
	DomainName *string `json:"domain_name,omitempty"`
}

func (o VersionHistoryOperator) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionHistoryOperator struct{}"
	}

	return strings.Join([]string{"VersionHistoryOperator", string(data)}, " ")
}
