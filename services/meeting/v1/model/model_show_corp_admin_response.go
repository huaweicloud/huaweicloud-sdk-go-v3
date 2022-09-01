package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCorpAdminResponse struct {

	// 用户id
	Id *string `json:"id,omitempty" xml:"id"`

	// 用户账号
	Account *string `json:"account,omitempty" xml:"account"`

	// 名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 管理员类型。 0：默认管理员 1：普通管理员
	AdminType *int32 `json:"adminType,omitempty" xml:"adminType"`

	// 邮箱
	Email *string `json:"email,omitempty" xml:"email"`

	// 手机号
	Phone *string `json:"phone,omitempty" xml:"phone"`

	// 手机号所属的国家
	Country *string `json:"country,omitempty" xml:"country"`

	Dept           *DeptBasicDto `json:"dept,omitempty" xml:"dept"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowCorpAdminResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCorpAdminResponse struct{}"
	}

	return strings.Join([]string{"ShowCorpAdminResponse", string(data)}, " ")
}
