package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库用户信息。
type ListGaussMySqlDatabaseUser struct {

	// 数据库用户名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 主机地址。
	Host *string `json:"host,omitempty" xml:"host"`

	// 数据库列表。
	Databases *[]ListGaussMySqlDatabase `json:"databases,omitempty" xml:"databases"`
}

func (o ListGaussMySqlDatabaseUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlDatabaseUser struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlDatabaseUser", string(data)}, " ")
}
