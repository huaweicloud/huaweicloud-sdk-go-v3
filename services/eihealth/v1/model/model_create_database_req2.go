package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseReq2 创建药物数据库请求体
type CreateDatabaseReq2 struct {

	// 数据库名称，长度为5-32个字符，首位需以小写英文字母开头，仅可以使用小写字母、数字、下划线“_”和中划线“-”
	Name string `json:"name"`

	// 数据库描述
	Description *string `json:"description,omitempty"`

	// css集群id
	CssId string `json:"css_id"`

	File *DatabaseFile `json:"file"`

	// 数据文件列名
	Columns []string `json:"columns"`

	// 是否打开组织共享
	Shareable *bool `json:"shareable,omitempty"`
}

func (o CreateDatabaseReq2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseReq2 struct{}"
	}

	return strings.Join([]string{"CreateDatabaseReq2", string(data)}, " ")
}
