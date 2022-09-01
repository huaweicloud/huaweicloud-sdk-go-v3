package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type KeystoneUserResult struct {

	// IAM用户密码状态。true：需要修改密码，false：正常。
	PwdStatus *bool `json:"pwd_status,omitempty" xml:"pwd_status"`

	// IAM用户所属账号ID。
	DomainId string `json:"domain_id" xml:"domain_id"`

	// IAM用户退出系统前，在控制台最后访问的项目ID。
	LastProjectId *string `json:"last_project_id,omitempty" xml:"last_project_id"`

	// IAM用户名。
	Name string `json:"name" xml:"name"`

	// IAM用户描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// IAM用户密码过期时间（UTC时间），“null”表示密码不过期。
	PasswordExpiresAt string `json:"password_expires_at" xml:"password_expires_at"`

	Links *Links `json:"links" xml:"links"`

	// IAM用户ID。
	Id string `json:"id" xml:"id"`

	// IAM用户是否启用。true表示启用，false表示停用，默认为true。
	Enabled bool `json:"enabled" xml:"enabled"`

	// IAM用户的密码强度。high：密码强度高；mid：密码强度中等；low：密码强度低。
	PwdStrength *string `json:"pwd_strength,omitempty" xml:"pwd_strength"`

	Extra *KeystoneUserResultExtra `json:"extra,omitempty" xml:"extra"`
}

func (o KeystoneUserResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUserResult struct{}"
	}

	return strings.Join([]string{"KeystoneUserResult", string(data)}, " ")
}
