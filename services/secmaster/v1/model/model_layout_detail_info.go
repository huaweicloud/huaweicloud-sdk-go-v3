package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LayoutDetailInfo 布局详情
type LayoutDetailInfo struct {

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

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建者ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 父布局ID
	ParentId *string `json:"parent_id,omitempty"`

	// 创建者名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 英文描述信息
	EnDescription *string `json:"en_description,omitempty"`

	// 订阅包id
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 英文名称
	EnName *string `json:"en_name,omitempty"`

	// 布局下的wizard id列表
	LayoutJson *string `json:"layout_json,omitempty"`

	// 租户ID
	ProjectId *string `json:"project_id,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 用户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 模板缩略图，当布局为模板时不为空
	Thumbnail *string `json:"thumbnail,omitempty"`

	// 被什么业务使用，DATACLASS/AOP_WORKFLOW/SECURITY_REPORT/DASHBOARD，及对应的字段
	UsedBy *string `json:"used_by,omitempty"`

	// 前端根据该值绑定图标
	LayoutCfg *string `json:"layout_cfg,omitempty"`

	// 布局类型；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	LayoutType *string `json:"layout_type,omitempty"`

	// 数据类ID或流程ID；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	BindingId *string `json:"binding_id,omitempty"`

	// 数据类名称或流程名称；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	BindingName *string `json:"binding_name,omitempty"`

	// 数据类或流程英文名称；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	BindingCode *string `json:"binding_code,omitempty"`

	// 字段总数；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	FieldsSum *int32 `json:"fields_sum,omitempty"`

	// 页面总数；used_by为SECURITY_REPORT/DASHBOARD时不返回该字段
	WizardsSum *int32 `json:"wizards_sum,omitempty"`

	// 系统区块总数
	SectionsSum *int32 `json:"sections_sum,omitempty"`

	// 系统模块总数
	ModulesSum *int32 `json:"modules_sum,omitempty"`

	// 自定义指标总数
	TabsSum *int32 `json:"tabs_sum,omitempty"`

	// 安全云脑版本
	Version *string `json:"version,omitempty"`

	// BOA底座版本
	BoaVersion *string `json:"boa_version,omitempty"`
}

func (o LayoutDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayoutDetailInfo struct{}"
	}

	return strings.Join([]string{"LayoutDetailInfo", string(data)}, " ")
}
