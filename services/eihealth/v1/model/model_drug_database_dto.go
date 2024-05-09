package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DrugDatabaseDto struct {

	// 数据库id
	Id *string `json:"id,omitempty"`

	// 数据库名称
	Name *string `json:"name,omitempty"`

	// 数据库类型
	Type *string `json:"type,omitempty"`

	// 数据库状态
	Status *string `json:"status,omitempty"`

	// 数据库描述
	Description *string `json:"description,omitempty"`

	// 数据库创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 数据库更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 创建数据库的用户名称
	Creator *string `json:"creator,omitempty"`

	// 失败提示，当作业执行失败时会返回
	FailedMessage *string `json:"failed_message,omitempty"`

	// css集群id
	CssId *string `json:"css_id,omitempty"`

	// css集群名称
	CssName *string `json:"css_name,omitempty"`

	// 数据库文件列表
	Files *[]DetailDatabaseFile `json:"files,omitempty"`

	// 数据库列名
	Columns *[]string `json:"columns,omitempty"`

	// 是否打开组织共享
	Shareable *bool `json:"shareable,omitempty"`

	// 分子数量
	DataNum *int32 `json:"data_num,omitempty"`
}

func (o DrugDatabaseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DrugDatabaseDto struct{}"
	}

	return strings.Join([]string{"DrugDatabaseDto", string(data)}, " ")
}
