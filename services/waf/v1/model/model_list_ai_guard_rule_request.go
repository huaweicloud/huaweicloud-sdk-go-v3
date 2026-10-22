package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAiGuardRuleRequest Request Object
type ListAiGuardRuleRequest struct {

	// **参数解释：** 策略ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 最大数量 **约束限制：** 不涉及 **取值范围：** 1-1000 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 0-10000 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 规则名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 规则状态 **约束限制：** 不涉及 **取值范围：** enable, disable **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释：** 检测选项类型 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DetectOpt *string `json:"detect_opt,omitempty"`

	// **参数解释：** 检测类型 **约束限制：** 不涉及 **取值范围：** text(文本检测), image(图片检测) **默认取值：** 不涉及
	DetectType *ListAiGuardRuleRequestDetectType `json:"detect_type,omitempty"`

	// **参数解释：** 是否查询模板规则 **约束限制：** 不涉及 **取值范围：** true, false **默认取值：** 不涉及
	WithTemplate *bool `json:"with_template,omitempty"`
}

func (o ListAiGuardRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiGuardRuleRequest struct{}"
	}

	return strings.Join([]string{"ListAiGuardRuleRequest", string(data)}, " ")
}

type ListAiGuardRuleRequestDetectType struct {
	value string
}

type ListAiGuardRuleRequestDetectTypeEnum struct {
	TEXT  ListAiGuardRuleRequestDetectType
	IMAGE ListAiGuardRuleRequestDetectType
}

func GetListAiGuardRuleRequestDetectTypeEnum() ListAiGuardRuleRequestDetectTypeEnum {
	return ListAiGuardRuleRequestDetectTypeEnum{
		TEXT: ListAiGuardRuleRequestDetectType{
			value: "text",
		},
		IMAGE: ListAiGuardRuleRequestDetectType{
			value: "image",
		},
	}
}

func (c ListAiGuardRuleRequestDetectType) Value() string {
	return c.value
}

func (c ListAiGuardRuleRequestDetectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAiGuardRuleRequestDetectType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
