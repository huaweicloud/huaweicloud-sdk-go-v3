package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StyleExtraMeta 风格额外信息
type StyleExtraMeta struct {

	// 是否支持照片建模
	PictureModelingEnable *bool `json:"picture_modeling_enable,omitempty"`

	// 是否支持模型编辑
	EditEnable *bool `json:"edit_enable,omitempty"`

	// 编辑使用引擎
	EditEngine *string `json:"edit_engine,omitempty"`

	// 值可设置条目列表
	EditValueItems map[string]StyleExtraMetaEditValueItems `json:"edit_value_items,omitempty"`

	// 颜色可设置条目列表
	EditColorItems map[string]StyleExtraMetaEditColorItems `json:"edit_color_items,omitempty"`

	// 可替换组件列表
	EditComponents map[string]StyleExtraMetaEditComponents `json:"edit_components,omitempty"`

	ModellingAlgorithm map[string]StyleExtraMetaModellingAlgorithm `json:"modelling_algorithm,omitempty"`
}

func (o StyleExtraMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StyleExtraMeta struct{}"
	}

	return strings.Join([]string{"StyleExtraMeta", string(data)}, " ")
}
