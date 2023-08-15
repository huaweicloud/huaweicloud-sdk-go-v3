package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigDataSourceView 数据连接结构体信息
type ApigDataSourceView struct {

	// 数据连接名称
	DwName *string `json:"dw_name,omitempty"`

	// 数据连接类型
	DwType *string `json:"dw_type,omitempty"`

	// 数据连接创建者
	CreateUser *string `json:"create_user,omitempty"`

	// 数据连接创建时间，时间戳
	CreateTime float32 `json:"create_time,omitempty"`

	// 代理id
	AgentId *string `json:"agent_id,omitempty"`

	// 代理名称
	AgentName *string `json:"agent_name,omitempty"`

	// 数据连接id
	DwId *string `json:"dw_id,omitempty"`

	// 数据连接限定名称
	QualifiedName *string `json:"qualified_name,omitempty"`

	// 数据连接描述
	Description *string `json:"description,omitempty"`
}

func (o ApigDataSourceView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigDataSourceView struct{}"
	}

	return strings.Join([]string{"ApigDataSourceView", string(data)}, " ")
}
