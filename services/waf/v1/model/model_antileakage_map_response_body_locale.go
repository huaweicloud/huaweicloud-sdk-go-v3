package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AntileakageMapResponseBodyLocale **参数解释：** 敏感信息选项里，各个类型的描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type AntileakageMapResponseBodyLocale struct {

	// **参数解释：** 响应码拦截，用于捕获和处理特定HTTP响应码的机制 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释：** 身份证号，用于识别个人身份的唯一编码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	IdCard *string `json:"id_card,omitempty"`

	// **参数解释：** 敏感信息过滤，用于检测和处理敏感信息的模块 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Sensitive *string `json:"sensitive,omitempty"`

	// **参数解释：** 电话号码，用于联系的数字编码 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Phone *string `json:"phone,omitempty"`

	// **参数解释：** 选项涉及的各种响应码，示例值400，401,404 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ResponseCode *string `json:"responseCode,omitempty"`

	// **参数解释：** 电子邮箱，用于电子通信的地址 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Email *string `json:"email,omitempty"`
}

func (o AntileakageMapResponseBodyLocale) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntileakageMapResponseBodyLocale struct{}"
	}

	return strings.Join([]string{"AntileakageMapResponseBodyLocale", string(data)}, " ")
}
