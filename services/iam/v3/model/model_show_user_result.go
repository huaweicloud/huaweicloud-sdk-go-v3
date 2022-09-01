package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ShowUserResult struct {

	// IAM用户是否启用。true表示启用，false表示停用，默认为true。
	Enabled bool `json:"enabled" xml:"enabled"`

	// IAM用户ID。
	Id string `json:"id" xml:"id"`

	// IAM用户所属账号ID。
	DomainId string `json:"domain_id" xml:"domain_id"`

	// IAM用户名。
	Name string `json:"name" xml:"name"`

	Links *Links `json:"links" xml:"links"`

	// IAM用户在外部系统中的ID。
	XuserId *string `json:"xuser_id,omitempty" xml:"xuser_id"`

	// IAM用户在外部系统中的类型。
	XuserType *string `json:"xuser_type,omitempty" xml:"xuser_type"`

	// IAM用户手机号的国家码。
	Areacode *string `json:"areacode,omitempty" xml:"areacode"`

	// IAM用户邮箱。
	Email *string `json:"email,omitempty" xml:"email"`

	// IAM用户手机号。
	Phone *string `json:"phone,omitempty" xml:"phone"`

	// IAM用户密码状态。true：需要修改密码，false：正常。
	PwdStatus *bool `json:"pwd_status,omitempty" xml:"pwd_status"`

	// IAM用户更新时间。
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// IAM用户创建时间。
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// IAM用户最后登录时间。
	LastLoginTime *string `json:"last_login_time,omitempty" xml:"last_login_time"`

	// IAM用户密码强度。结果为Low/Middle/High/None，分别表示密码强度低/中/高/无。
	PwdStrength *string `json:"pwd_strength,omitempty" xml:"pwd_strength"`

	// IAM用户是否为根用户。
	IsDomainOwner bool `json:"is_domain_owner" xml:"is_domain_owner"`

	// IAM用户访问模式。
	AccessMode string `json:"access_mode" xml:"access_mode"`

	// IAM用户描述信息
	Description string `json:"description" xml:"description"`
}

func (o ShowUserResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserResult struct{}"
	}

	return strings.Join([]string{"ShowUserResult", string(data)}, " ")
}
