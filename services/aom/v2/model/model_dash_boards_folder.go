package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DashBoardsFolder 仪表盘分组列表。
type DashBoardsFolder struct {

	// 仪表盘分组创建时间。
	Created int64 `json:"created"`

	// 仪表盘分组更新时间。
	Updated int64 `json:"updated"`

	// 仪表盘分组创建账号。
	CreatedBy string `json:"created_by"`

	// 仪表盘分组下仪表盘id列表。
	DashboardIds []string `json:"dashboard_ids"`

	// 是否展示仪表盘分组。
	Display bool `json:"display"`

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 仪表盘分组id。
	FolderId string `json:"folder_id"`

	// 仪表盘分组名称。
	FolderTitle string `json:"folder_title"`

	// 仪表盘分组英文名称。
	FolderTitleEn *string `json:"folder_title_en,omitempty"`

	// 仪表盘分组是否为默认仪表盘分组。
	IsTemplate bool `json:"is_template"`
}

func (o DashBoardsFolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DashBoardsFolder struct{}"
	}

	return strings.Join([]string{"DashBoardsFolder", string(data)}, " ")
}
