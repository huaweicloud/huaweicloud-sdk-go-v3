package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LayoutUpdateRequestPojo struct {

	// 布局名称
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 订阅包id
	CloudPackId *string `json:"cloud_pack_id,omitempty"`

	// 订阅包名称
	CloudPackName *string `json:"cloud_pack_name,omitempty"`

	// 订阅包版本
	CloudPackVersion *string `json:"cloud_pack_version,omitempty"`

	// 是否内置，true内置，false非内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 是否为模板
	IsTemplate *bool `json:"is_template,omitempty"`

	// 布局信息
	LayoutJson *interface{} `json:"layout_json,omitempty"`

	// 字段总数；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	FieldsSum *int32 `json:"fields_sum,omitempty"`

	// 页面总数；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	WizardsSum *int32 `json:"wizards_sum,omitempty"`

	// 系统指标总数
	SectionsSum *int32 `json:"sections_sum,omitempty"`

	// 自定义指标总数
	TabsSum *int32 `json:"tabs_sum,omitempty"`

	// BOA底座版本
	BoaVersion *string `json:"boa_version,omitempty"`
}

func (o LayoutUpdateRequestPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayoutUpdateRequestPojo struct{}"
	}

	return strings.Join([]string{"LayoutUpdateRequestPojo", string(data)}, " ")
}
