package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库信息。
type ListGaussMysqlDatabaseInfo struct {

	// 数据库名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 数据库使用的字符集，如utf8mb4、gbk等。
	Charset *string `json:"charset,omitempty" xml:"charset"`

	// 已授权数据库用户列表。
	Users *[]GaussMySqlDatabaseInfo `json:"users,omitempty" xml:"users"`
}

func (o ListGaussMysqlDatabaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMysqlDatabaseInfo struct{}"
	}

	return strings.Join([]string{"ListGaussMysqlDatabaseInfo", string(data)}, " ")
}
