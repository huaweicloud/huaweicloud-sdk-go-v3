package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Connection 数据连接
type Connection struct {

	// 关联guid
	Guid *string `json:"guid,omitempty"`

	// 显示内容
	DisplayText *string `json:"display_text,omitempty"`

	// 类型名称
	TypeName *string `json:"type_name,omitempty"`

	// 连接类型
	ConnectionType *string `json:"connection_type,omitempty"`

	// 限定名称
	QualifiedName *string `json:"qualified_name,omitempty"`
}

func (o Connection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Connection struct{}"
	}

	return strings.Join([]string{"Connection", string(data)}, " ")
}
