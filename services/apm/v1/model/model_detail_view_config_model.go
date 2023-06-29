package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetailViewConfigModel 详情视图配置。
type DetailViewConfigModel struct {

	// 视图对应的指标集的名称。
	MetricSet *string `json:"metric_set,omitempty"`

	// 过滤参数。
	FilterPrefix *string `json:"filter_prefix,omitempty"`

	// 视图函数集合。
	DetailViewItemList *[]DetailViewItem `json:"detail_view_item_list,omitempty"`
}

func (o DetailViewConfigModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetailViewConfigModel struct{}"
	}

	return strings.Join([]string{"DetailViewConfigModel", string(data)}, " ")
}
