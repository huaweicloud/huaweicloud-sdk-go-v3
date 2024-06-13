package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MindmapPageParamV3 struct {

	// 目录ID集合
	FolderIdCollection *[]string `json:"folder_id_collection,omitempty"`

	// 创建者ID集合
	CreatorNumCollection *[]string `json:"creator_num_collection,omitempty"`

	// 更新人ID集合
	UpdaterNumCollection *[]string `json:"updater_num_collection,omitempty"`

	// 根目录ID
	FolderRootId *string `json:"folder_root_id,omitempty"`

	// 主键ID集合
	IdCollection *[]string `json:"id_collection,omitempty"`

	// 起始偏移量，表示从此偏移量开始查询，offset大于等于0，小于等于100000
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，最大支持200条
	Limit *int32 `json:"limit,omitempty"`

	// 脑图名称
	Name *string `json:"name,omitempty"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 分支uri
	BranchUri *string `json:"branch_uri,omitempty"`

	// 是否基线
	IsMaster *int32 `json:"is_master,omitempty"`

	// 计划uri
	IteratorUri *string `json:"iterator_uri,omitempty"`
}

func (o MindmapPageParamV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MindmapPageParamV3 struct{}"
	}

	return strings.Join([]string{"MindmapPageParamV3", string(data)}, " ")
}
