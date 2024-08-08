package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWarehouseAppsRequest Request Object
type ListWarehouseAppsRequest struct {

	// 单次查询的大小[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 审核状态。
	VerifyStatus *string `json:"verify_status,omitempty"`

	// 应用仓库中的应用记录ID。
	AppId *string `json:"app_id,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 应用分类： * `GAME`-  `游戏`。 * `BUSSINESS_INTELLIGENCE`- `商业智能`。 * `SECURE_STORAGE`-  `安全与存储`。 * `MULTIMEDIA_AND_CODING`- `多媒体与编解码`。 * `PROJECT_MANAGEMENT`- `项目管理`, * `PRODUCTIVITY_AND_COLLABORATION`-  `生产力与协作`。 * `WEB_ADN_APPLICATION`-  `网页与应用开发`。 * `GRAPHIC_DESIGN`-  `图形设计`。 * `OTHER`-  `其它`。
	AppCategory *string `json:"app_category,omitempty"`
}

func (o ListWarehouseAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWarehouseAppsRequest struct{}"
	}

	return strings.Join([]string{"ListWarehouseAppsRequest", string(data)}, " ")
}
