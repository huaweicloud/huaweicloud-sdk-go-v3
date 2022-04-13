package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CreateUserResult struct {
	// IAM用户访问方式。 - default：默认访问模式，编程访问和管理控制台访问。 - programmatic：编程访问。 - console：管理控制台访问。

	AccessMode *string `json:"access_mode,omitempty"`
	// IAM用户状态信息。

	Status *int32 `json:"status,omitempty"`
	// IAM用户首次登录是否重置密码。

	PwdStatus *bool `json:"pwd_status,omitempty"`
	// IAM用户在外部系统中的ID。 >外部系统指与华为云对接的外部企业管理系统，xaccount_type、xaccount_id、xdomain_type、xdomain_id、xuser_type、xuser_id等参数值，无法在华为云获取，请咨询企业管理员。

	XuserId *string `json:"xuser_id,omitempty"`
	// 用户在外部系统中的类型。 >外部系统指与华为云对接的外部企业管理系统，xaccount_type、xaccount_id、xdomain_type、xdomain_id、xuser_type、xuser_id等参数值，无法在华为云获取，请咨询企业管理员。

	XuserType *string `json:"xuser_type,omitempty"`
	// IAM用户描述信息。

	Description *string `json:"description,omitempty"`
	// IAM用户名，长度5~32字符之间，首位不能为数字，特殊字符只能包含下划线“_”、中划线“-”和空格。

	Name string `json:"name"`
	// IAM用户手机号，纯数字，长度小于等于32字符。必须与国家码同时存在。

	Phone *string `json:"phone,omitempty"`
	// IAM用户是否为账号管理员。

	IsDomainOwner *bool `json:"is_domain_owner,omitempty"`
	// IAM用户所属账号ID。

	DomainId string `json:"domain_id"`
	// 是否启用IAM用户。true为启用，false为停用，默认为true。

	Enabled bool `json:"enabled"`
	// 国家码。中国大陆为“0086”。

	Areacode *string `json:"areacode,omitempty"`
	// IAM用户邮箱。

	Email *string `json:"email,omitempty"`
	// IAM用户创建时间。

	CreateTime *string `json:"create_time,omitempty"`
	// 运营主体的客户编码。

	XdomainId *string `json:"xdomain_id,omitempty"`
	// 运营主体。

	XdomainType *string `json:"xdomain_type,omitempty"`
	// IAM用户ID。

	Id string `json:"id"`
	// 密码过期时间（UTC时间），“null”表示密码不过期。

	PasswordExpiresAt *string `json:"password_expires_at,omitempty"`
}

func (o CreateUserResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserResult struct{}"
	}

	return strings.Join([]string{"CreateUserResult", string(data)}, " ")
}
