package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDashBoardResponse Response Object
type ShowDashBoardResponse struct {

	// 项目ID，可以从控制台获取，也可以从调用API处获取。获取方式请参见：[获取项目ID](aom_04_0024.xml)。
	ProjectId *string `json:"project_id,omitempty"`

	// 仪表盘类型。
	DashboardType *string `json:"dashboard_type,omitempty"`

	// 仪表盘名称。
	DashboardTitle *string `json:"dashboard_title,omitempty"`

	// 仪表盘英文名称。
	DashboardTitleEn *string `json:"dashboard_title_en,omitempty"`

	// 仪表盘id。
	DashboardId *string `json:"dashboard_id,omitempty"`

	// 仪表盘版本。
	Version *string `json:"version,omitempty"`

	// 仪表盘企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 仪表盘分组名称。
	FolderName *string `json:"folder_name,omitempty"`

	// 仪表盘分组id。
	FolderId *string `json:"folder_id,omitempty"`

	// 待同步的仪表盘数。
	SyncData *string `json:"sync_data,omitempty"`

	// 是否创建 - false：更新 - true：创建
	IsCreateAction *bool `json:"is_create_action,omitempty"`

	// 仪表盘标签列表。
	DashboardTags *[]map[string]string `json:"dashboard_tags,omitempty"`

	// 是否收藏 - true：收藏 - false：不收藏
	IsFavorite *bool `json:"is_favorite,omitempty"`

	// 仪表盘创建时间。
	Created *int64 `json:"created,omitempty"`

	// 仪表盘更新时间。
	Updated *int64 `json:"updated,omitempty"`

	// 创建仪表盘的账号名称。
	CreatedBy *string `json:"created_by,omitempty"`

	// 更新仪表盘的账号名称。
	UpdatedBy *string `json:"updated_by,omitempty"`

	// 仪表盘图表详情。
	Charts *interface{} `json:"charts,omitempty"`

	// 仪表盘变量列表。
	Templating *interface{} `json:"templating,omitempty"`

	// 是否展示。
	Display *bool `json:"display,omitempty"`

	// 查询总次数。
	QueryCount *int32 `json:"query_count,omitempty"`

	// 默认查询时间范围。
	TimeRange      *string `json:"time_range,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDashBoardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDashBoardResponse struct{}"
	}

	return strings.Join([]string{"ShowDashBoardResponse", string(data)}, " ")
}
