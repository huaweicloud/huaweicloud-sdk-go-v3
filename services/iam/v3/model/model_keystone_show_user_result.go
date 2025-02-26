package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowUserResult
type KeystoneShowUserResult struct {

	// IAM用户密码状态。true：需要修改密码，false：正常。
	PwdStatus *bool `json:"pwd_status,omitempty"`

	// IAM用户所属账号ID。
	DomainId string `json:"domain_id"`

	// IAM用户退出系统前，在控制台最后访问的项目ID。
	LastProjectId *string `json:"last_project_id,omitempty"`

	// IAM用户名。
	Name string `json:"name"`

	// IAM用户描述信息。
	Description *string `json:"description,omitempty"`

	// IAM用户密码过期时间（UTC时间），“null”表示密码不过期。
	PasswordExpiresAt string `json:"password_expires_at"`

	Links *Links `json:"links"`

	// IAM用户ID。
	Id string `json:"id"`

	// IAM用户是否启用。true表示启用，false表示停用，默认为true。
	Enabled bool `json:"enabled"`

	// IAM用户访问方式。 - default：默认访问模式，编程访问和管理控制台访问。 - programmatic：编程访问。 - console：管理控制台访问。
	AccessMode *string `json:"access_mode,omitempty"`
}

func (o KeystoneShowUserResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowUserResult struct{}"
	}

	return strings.Join([]string{"KeystoneShowUserResult", string(data)}, " ")
}
