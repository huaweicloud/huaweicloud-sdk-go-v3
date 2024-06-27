package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MindmapObject struct {

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 创建人工号
	CreatorNum *string `json:"creator_num,omitempty"`

	// 文件路径
	FolderId *string `json:"folder_id,omitempty"`

	// 根目录id
	FolderRootId *string `json:"folder_root_id,omitempty"`

	// id 主键
	Id *string `json:"id,omitempty"`

	// 脑图最大深度
	MaxDepth *int32 `json:"max_depth,omitempty"`

	// 脑图JSON
	Mindmap *string `json:"mindmap,omitempty"`

	// 脑图名称
	Name *string `json:"name,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o MindmapObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MindmapObject struct{}"
	}

	return strings.Join([]string{"MindmapObject", string(data)}, " ")
}
