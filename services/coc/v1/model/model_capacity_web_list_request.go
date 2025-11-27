package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CapacityWebListRequest **参数解释：** 应用容量数据请求体。 **约束限制：** 不涉及。 **取值范围：** 请求体内容按照properties下的参数说明传参。 **默认取值：** 不涉及。
type CapacityWebListRequest struct {

	// **参数解释：** 用户选择当前分组对应的id值。 **约束限制：** 应用、组件和分组ID，有且仅有1个有值。 **取值范围：** 字符串，长度24个字符。 **默认取值：** 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释：** 用户选择当前组件对应的id值。 **约束限制：** 应用、组件和分组ID，有且仅有1个有值。 **取值范围：** 字符串，长度24个字符。 **默认取值：** 不涉及。
	ComponentId *string `json:"component_id,omitempty"`

	// **参数解释：** 用户选择当前应用对应的id值。 **约束限制：** 应用、组件和分组ID，有且仅有1个有值。 **取值范围：** 字符串，长度24个字符。 **默认取值：** 不涉及。
	ApplicationId *string `json:"application_id,omitempty"`

	// **参数解释：** 用户登录租户对应的账号ID即租户id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 资源对象。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ProviderObj []CapacityWebListRequestProviderObj `json:"provider_obj"`
}

func (o CapacityWebListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapacityWebListRequest struct{}"
	}

	return strings.Join([]string{"CapacityWebListRequest", string(data)}, " ")
}
