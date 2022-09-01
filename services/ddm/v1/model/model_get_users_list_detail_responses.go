package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto query response Object
type GetUsersListDetailResponses struct {

	// DDM实例帐号名称。
	Name string `json:"name" xml:"name"`

	// DDM实例帐号状态。
	Status string `json:"status" xml:"status"`

	// DDM实例帐号的基础权限。  取值为：CREATE、DROP、ALTER、INDEX、INSERT、DELETE、UPDATE、SELECT
	BaseAuthority []string `json:"base_authority" xml:"base_authority"`

	// DDM实例帐号的扩展权限。2021年8月开始不支持该字段，9月会去掉该字段。  取值为：fulltableDelete、fulltableSelect、fulltableUpdate
	ExtendAuthority *[]string `json:"extend_authority,omitempty" xml:"extend_authority"`

	// DDM实例帐号的描述。
	Description string `json:"description" xml:"description"`

	// DDM实例帐号的创建时间。
	Created int64 `json:"created" xml:"created"`

	// 关联的逻辑库的集合。
	Databases []GetUsersListdatabase `json:"databases" xml:"databases"`
}

func (o GetUsersListDetailResponses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetUsersListDetailResponses struct{}"
	}

	return strings.Join([]string{"GetUsersListDetailResponses", string(data)}, " ")
}
