package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLdapConfigResponse Response Object
type ShowLdapConfigResponse struct {

	// ldap服务器的url
	Url *string `json:"url,omitempty"`

	// 数据库中的域
	BaseDn *string `json:"base_dn,omitempty"`

	// 用户区别名
	UserDn *string `json:"user_dn,omitempty"`

	// 过滤条件。保留字段，暂不支持
	FilterCondition *string `json:"filter_condition,omitempty"`

	// ldap备节点的url
	BackupUrl *string `json:"backup_url,omitempty"`

	// ldap的schema，不填写则默认为RFC2307
	Schema *string `json:"schema,omitempty"`

	// ldap搜索的超时时间，单位为秒。不填写则默认为3秒
	SearchTimeout *int32 `json:"search_timeout,omitempty"`

	// 访问ldap服务器失败后是否允许使用本地用户鉴权
	AllowLocalUser *string `json:"allow_local_user,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLdapConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLdapConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowLdapConfigResponse", string(data)}, " ")
}
