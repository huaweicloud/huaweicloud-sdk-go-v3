package model

import (
	"encoding/json"

	"strings"
)

// This is a auto query response Object
type GetUsersListDetailResponses struct {
	// DDM实例帐号名称。

	Name string `json:"name"`
	// DDM实例帐号状态。

	Status string `json:"status"`
	// DDM实例帐号的基础权限。  取值为：CREATE、DROP、ALTER、INDEX、INSERT、DELETE、UPDATE、SELECT

	BaseAuthority []string `json:"base_authority"`
	// DDM实例帐号的扩展权限。2021年8月开始不支持该字段，9月会去掉该字段。  取值为：fulltableDelete、fulltableSelect、fulltableUpdate

	ExtendAuthority *[]string `json:"extend_authority,omitempty"`
	// DDM实例帐号的描述。

	Description string `json:"description"`
	// DDM实例帐号的创建时间。

	Created int64 `json:"created"`
	// 关联的逻辑库的集合。

	Databases []GetUsersListdatabase `json:"databases"`
}

func (o GetUsersListDetailResponses) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "GetUsersListDetailResponses struct{}"
	}

	return strings.Join([]string{"GetUsersListDetailResponses", string(data)}, " ")
}
