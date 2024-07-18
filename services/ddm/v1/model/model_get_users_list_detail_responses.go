package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetUsersListDetailResponses This is a auto query response Object
type GetUsersListDetailResponses struct {

	// DDM实例帐号名称。
	Name string `json:"name"`

	// DDM实例帐号状态。
	Status string `json:"status"`

	// DDM实例帐号的基础权限。  取值为：CREATE、DROP、ALTER、INDEX、INSERT、DELETE、UPDATE、SELECT
	BaseAuthority []string `json:"base_authority"`

	// DDM实例账号的密码修改时间，UNIX时间戳格式。
	PasswordLastChanged *int64 `json:"password_last_changed,omitempty"`

	// DDM实例帐号的扩展权限。2021年8月开始不支持该字段，9月会去掉该字段。  取值为：fulltableDelete、fulltableSelect、fulltableUpdate
	ExtendAuthority *[]string `json:"extend_authority,omitempty"`

	// DDM实例帐号的描述。
	Description string `json:"description"`

	// DDM实例帐号的创建时间，UNIX时间戳格式。
	Created int64 `json:"created"`

	// 关联的逻辑库的集合。
	Databases []GetUsersListdatabase `json:"databases"`
}

func (o GetUsersListDetailResponses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetUsersListDetailResponses struct{}"
	}

	return strings.Join([]string{"GetUsersListDetailResponses", string(data)}, " ")
}
