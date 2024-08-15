package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimTemplate 短信下发响应对象。
type AimTemplate struct {

	// 智能信息模板ID。  > 智能信息平台生成的模板ID，由9位数字组成。
	TplId *string `json:"tpl_id,omitempty"`

	// 智能信息模板名称。
	TplName *string `json:"tpl_name,omitempty"`

	// 场景类型。
	Scene *string `json:"scene,omitempty"`

	// 模板状态。  - 0：禁用 - 1：启用
	TplState *int32 `json:"tpl_state,omitempty"`

	// 禁用原因。
	DisableDesc *string `json:"disable_desc,omitempty"`

	// 禁用时间。样例：1970-01-01T00:00 :00Z。
	DisableTime *string `json:"disable_time,omitempty"`

	// 审核状态。 - 0：未提交  - 1：审核中  - 2：审核通过  - 3：审核失败
	AuditState *int32 `json:"audit_state,omitempty"`

	// 审批信息。
	AuditDesc *string `json:"audit_desc,omitempty"`

	// 短信示例。  >对应创建个人模板API中的入参sms_example。
	Description *string `json:"description,omitempty"`

	// 创建时间。样例为：1970-01-01T00:00:00Z。
	CreationTime *string `json:"creation_time,omitempty"`

	// 更新时间。样例为：1970-01-01T00:00:00Z。
	UpdateTime *string `json:"update_time,omitempty"`

	// 模板页面HTML，JSON格式。
	Pages *string `json:"pages,omitempty"`

	// 模板动态参数列表。
	Params *[]AimTemplateParams `json:"params,omitempty"`

	// 支持厂商列表。
	FactoryInfo *[]FactoryInfo `json:"factory_info,omitempty"`

	// 审核状态。 - 1：短链解析模板 - 2：文本识别模板 - 4：一体化模板
	MatchType *int32 `json:"match_type,omitempty"`

	// 布局类型。
	CardId *string `json:"card_id,omitempty"`

	// sub_type。
	SubType *int32 `json:"sub_type,omitempty"`

	// 模板二维码预览地址。
	PreviewUrl *string `json:"preview_url,omitempty"`
}

func (o AimTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimTemplate struct{}"
	}

	return strings.Join([]string{"AimTemplate", string(data)}, " ")
}
