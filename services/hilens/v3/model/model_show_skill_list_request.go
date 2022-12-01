package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSkillListRequest struct {

	// 每页显示的条目数量, 最大 100，默认值 10
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始位置, 默认值 0
	Offset *int32 `json:"offset,omitempty"`

	// 技能名称，支持模糊匹配。中英文、数字、下划线、中划线 长度[1-60]
	SkillName *string `json:"skill_name,omitempty"`

	// 技能形式，no_termplate不支持Modelbox部署模板，support_template支持Modelbox模板。
	SkillForm *string `json:"skill_form,omitempty"`

	// 技能可见权限，支持公开可见public以及白名单whitelist
	Permission *string `json:"permission,omitempty"`

	// 技能来源，分别hilens，ma_pro，studio
	TemplateSource *string `json:"template_source,omitempty"`

	// 技能审核状态状态，1表示审核通过，2表示审核不通过，0表示待审核
	Status *int32 `json:"status,omitempty"`

	// 收费模式，0表示永久免费，1表示收费，2表示30天试用，3表示365天试用，4表示收费（永久使用）
	ChargeModel *int32 `json:"charge_model,omitempty"`

	// 技能操作系统平台，其值为：Linux，Android， iOS， LiteOS，Windows
	Platform *string `json:"platform,omitempty"`

	// 技能芯片类型，其值为Ascend 310,Ascend 310(Atlas 200 DK)，Arm，x86，3516CV500,3519AV100,3519V101,3516DV300,3516EV200,3516EV300,3518EV300
	Chip *string `json:"chip,omitempty"`

	// 技能类型，lite表示使用于海思芯片的轻量型技能。standard表示标准技能。
	Type *string `json:"type,omitempty"`

	// 收费模式多选，0表示永久免费，1表示收费，2表示30天试用，3表示365天试用，4表示收费（永久使用），分隔符为|，例如输入为1|2|3
	ChargeModels *string `json:"charge_models,omitempty"`

	// 设备类型校验，允许输入多个设备类型，用|分隔，例如设备a|设备b。每种设备类型不允许#~^$%&*<>{}[]'|字符，长度1到100。最多50个设备类型
	DeviceTypes *string `json:"device_types,omitempty"`

	// 技能适用场景，支持多选，分隔符|，例如场景a|场景b，每种场景不允许输入#~^$%&*<>{}\\\\[\\\\]'\\\\|等字符，长度100以内，最多20个设备类型。
	Scenes *string `json:"scenes,omitempty"`
}

func (o ShowSkillListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillListRequest struct{}"
	}

	return strings.Join([]string{"ShowSkillListRequest", string(data)}, " ")
}
