package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MindmapBackup struct {

	// 备份名称
	BakName *string `json:"bak_name,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建人工号
	CreatorNum *string `json:"creator_num,omitempty"`

	// id 主键
	Id *string `json:"id,omitempty"`

	// 备注
	Memo *string `json:"memo,omitempty"`

	// 脑图JSON
	Mindmap *string `json:"mindmap,omitempty"`

	// 脑图Id
	MindmapId *string `json:"mindmap_id,omitempty"`

	// 备份类型
	Type *string `json:"type,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o MindmapBackup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MindmapBackup struct{}"
	}

	return strings.Join([]string{"MindmapBackup", string(data)}, " ")
}
