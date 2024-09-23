package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccountReqBody CreateAccount 操作的请求体。
type CreateAccountReqBody struct {

	// 账号名称
	Name string `json:"name"`

	// 邮箱
	Email *string `json:"email,omitempty"`

	// 手机号码
	Phone *string `json:"phone,omitempty"`

	// 委托名称
	AgencyName *string `json:"agency_name,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 要绑定到新创建的账号的标签列表。
	Tags *[]TagDto `json:"tags,omitempty"`
}

func (o CreateAccountReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccountReqBody struct{}"
	}

	return strings.Join([]string{"CreateAccountReqBody", string(data)}, " ")
}
