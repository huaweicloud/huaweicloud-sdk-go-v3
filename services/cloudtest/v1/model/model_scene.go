package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Scene struct {

	// app
	App *string `json:"app,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建人工号
	CreatorNum *string `json:"creator_num,omitempty"`

	// 删除时间
	DeleteTime *string `json:"delete_time,omitempty"`

	// 是否删除
	Deleted *string `json:"deleted,omitempty"`

	// id 主键
	Id *string `json:"id,omitempty"`

	// 脑图id
	MindmapId *string `json:"mindmap_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 节点id
	NodeId *string `json:"node_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 状态类型
	StatusType *string `json:"status_type,omitempty"`

	// tc数量
	TcCounts *string `json:"tc_counts,omitempty"`

	// 测试用例
	TestCases *string `json:"test_cases,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o Scene) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Scene struct{}"
	}

	return strings.Join([]string{"Scene", string(data)}, " ")
}
