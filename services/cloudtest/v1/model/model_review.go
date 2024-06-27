package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Review struct {

	// 名称
	Charger *string `json:"charger,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建人
	Creator *string `json:"creator,omitempty"`

	// 是否删除
	Deleted *string `json:"deleted,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// id 主键
	Id *string `json:"id,omitempty"`

	// 脑图id
	MindmapId *string `json:"mindmap_id,omitempty"`

	// 节点id
	NodeId *string `json:"node_id,omitempty"`

	// 节点值
	NodeValue *string `json:"node_value,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o Review) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Review struct{}"
	}

	return strings.Join([]string{"Review", string(data)}, " ")
}
