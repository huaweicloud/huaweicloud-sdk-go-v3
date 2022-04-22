package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCorpAdminResponse struct {

	// 用户id
	Id *string `json:"id,omitempty"`

	// 用户账号
	Account *string `json:"account,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 管理员类型。 0：默认管理员 1：普通管理员
	AdminType *int32 `json:"adminType,omitempty"`

	// 邮箱
	Email *string `json:"email,omitempty"`

	// 手机号
	Phone *string `json:"phone,omitempty"`

	// 手机号所属的国家
	Country *string `json:"country,omitempty"`

	Dept           *DeptBasicDto `json:"dept,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowCorpAdminResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCorpAdminResponse struct{}"
	}

	return strings.Join([]string{"ShowCorpAdminResponse", string(data)}, " ")
}
