package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MindmapRecycle struct {

	// app
	App *string `json:"app,omitempty"`

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

	// 脑图版本
	MapVersion *string `json:"map_version,omitempty"`

	// 脑图JSON
	Mindmap *string `json:"mindmap,omitempty"`

	// 脑图名称
	Name *string `json:"name,omitempty"`

	// 脑图是否私有
	Privacy *string `json:"privacy,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o MindmapRecycle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MindmapRecycle struct{}"
	}

	return strings.Join([]string{"MindmapRecycle", string(data)}, " ")
}
