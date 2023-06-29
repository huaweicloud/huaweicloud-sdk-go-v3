package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseObjectInfo 数据库对象信息
type DatabaseObjectInfo struct {

	// type为database时，为库名；type为table或者view时，字段值参考示例
	Id *string `json:"id,omitempty"`

	// type为table或view时需要填写，为库名
	ParentId *string `json:"parent_id,omitempty"`

	// 类型。
	Type *string `json:"type,omitempty"`

	// 数据库对象名称，库名、表名、视图名
	Name *string `json:"name,omitempty"`

	// 别名，映射的新名称
	AliasName *string `json:"alias_name,omitempty"`
}

func (o DatabaseObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObjectInfo struct{}"
	}

	return strings.Join([]string{"DatabaseObjectInfo", string(data)}, " ")
}
