package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 指向线视图配置。
type LineViewConfigModel struct {

	// 视图对应的指标集的名称。
	MetricSet *string `json:"metric_set,omitempty"`

	// 过滤参数。
	FilterPrefix *string `json:"filter_prefix,omitempty"`

	// 视图函数集合。
	LineViewItemList *[]LineViewItem `json:"line_view_item_list,omitempty"`
}

func (o LineViewConfigModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineViewConfigModel struct{}"
	}

	return strings.Join([]string{"LineViewConfigModel", string(data)}, " ")
}
