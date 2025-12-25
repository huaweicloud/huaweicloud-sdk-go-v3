package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LayoutCreateRequestPojo struct {

	// 布局名称
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者
	CreatorName *string `json:"creator_name,omitempty"`

	// 订阅包id
	CloudPackId *string `json:"cloud_pack_id,omitempty"`

	// 订阅包名称
	CloudPackName *string `json:"cloud_pack_name,omitempty"`

	// 订阅包版本
	CloudPackVersion *string `json:"cloud_pack_version,omitempty"`

	// 是否内置，true内置，false非内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 布局信息
	LayoutJson *interface{} `json:"layout_json,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 用户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 模板缩略图，当布局为模板时不为空
	Thumbnail *string `json:"thumbnail,omitempty"`

	// 被什么业务使用，DATACLASS/AOP_WORKFLOW/SECURITY_REPORT/DASHBOARD，及对应的字段
	UsedBy string `json:"used_by"`

	// 布局类型；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	LayoutType *string `json:"layout_type,omitempty"`

	// 数据类ID或流程ID；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	BindingId *string `json:"binding_id,omitempty"`

	// 数据类名称business_code；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	BindingCode *string `json:"binding_code,omitempty"`

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

func (o LayoutCreateRequestPojo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayoutCreateRequestPojo struct{}"
	}

	return strings.Join([]string{"LayoutCreateRequestPojo", string(data)}, " ")
}
