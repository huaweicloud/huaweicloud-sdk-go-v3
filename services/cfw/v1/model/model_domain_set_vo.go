package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainSetVo struct {

	// **参数解释**： 域名组id **取值范围**： 不涉及
	SetId *string `json:"set_id,omitempty"`

	// **参数解释**： 域名组名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 域名组描述 **取值范围**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 域名组被规则引用次数 **取值范围**： 不涉及
	RefCount *int32 `json:"ref_count,omitempty"`

	// **参数解释**： 域名组类型 **取值范围**： - 0：应用域名组 - 1：网络域名组
	DomainSetType *int32 `json:"domain_set_type,omitempty"`

	// **参数解释**： 配置状态 **取值范围**： - -1：未配置态 - 0：配置失败 - 1：配置成功 - 2：配置中 - 3：正常 - 4：配置异常
	ConfigStatus *int32 `json:"config_status,omitempty"`

	// **参数解释**： 使用规则id列表 **取值范围**： 不涉及
	Rules *[]UseRuleVo `json:"rules,omitempty"`
}

func (o DomainSetVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainSetVo struct{}"
	}

	return strings.Join([]string{"DomainSetVo", string(data)}, " ")
}
