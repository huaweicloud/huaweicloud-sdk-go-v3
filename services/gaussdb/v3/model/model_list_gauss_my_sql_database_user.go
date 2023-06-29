package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGaussMySqlDatabaseUser 数据库用户信息。
type ListGaussMySqlDatabaseUser struct {

	// 数据库用户名。
	Name *string `json:"name,omitempty"`

	// 主机地址。
	Host *string `json:"host,omitempty"`

	// 数据库用户备注。
	Comment *string `json:"comment,omitempty"`

	// 数据库列表。
	Databases *[]ListGaussMySqlDatabase `json:"databases,omitempty"`
}

func (o ListGaussMySqlDatabaseUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlDatabaseUser struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlDatabaseUser", string(data)}, " ")
}
