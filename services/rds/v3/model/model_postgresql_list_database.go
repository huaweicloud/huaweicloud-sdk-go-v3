package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库信息。
type PostgresqlListDatabase struct {

	// 数据库名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 数据库所属用户。
	Owner *string `json:"owner,omitempty" xml:"owner"`

	// 数据库使用的字符集，例如UTF8。
	CharacterSet *string `json:"character_set,omitempty" xml:"character_set"`

	// 数据库排序集，例如en_US.UTF-8等。
	CollateSet *string `json:"collate_set,omitempty" xml:"collate_set"`

	// 数据库大小（单位：字节）。
	Size *int64 `json:"size,omitempty" xml:"size"`
}

func (o PostgresqlListDatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgresqlListDatabase struct{}"
	}

	return strings.Join([]string{"PostgresqlListDatabase", string(data)}, " ")
}
