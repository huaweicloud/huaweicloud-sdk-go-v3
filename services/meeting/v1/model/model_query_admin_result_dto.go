package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 管理员信息
type QueryAdminResultDto struct {
	// 用户id

	Id *string `json:"id,omitempty"`
	// 用户账号

	Account *string `json:"account,omitempty"`
	// 名称

	Name *string `json:"name,omitempty"`
	// 管理员类型 - 0：默认管理员 - 1：普通管理员

	AdminType *int32 `json:"adminType,omitempty"`
	// 邮箱

	Email *string `json:"email,omitempty"`
	// 联系电话

	Phone *string `json:"phone,omitempty"`
	// 联系电话所属的国家

	Country *string `json:"country,omitempty"`
}

func (o QueryAdminResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAdminResultDto struct{}"
	}

	return strings.Join([]string{"QueryAdminResultDto", string(data)}, " ")
}
