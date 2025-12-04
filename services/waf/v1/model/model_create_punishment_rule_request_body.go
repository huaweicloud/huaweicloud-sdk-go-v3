package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreatePunishmentRuleRequestBody struct {

	// **参数解释：** 攻击惩罚类别 **约束限制：** 不涉及 **取值范围：**  - long_ip_block  - long_cookie_block  - long_params_block  - short_ip_block  - short_cookie_block  - short_params_block  **默认取值：** 不涉及
	Category CreatePunishmentRuleRequestBodyCategory `json:"category"`

	// 拦截时间，如果选择前缀为long的攻击惩罚类别，则block_time时长范围设置为301-1800;选择前缀为short的攻击惩罚类别，则block_time时长范围为0-300之间
	BlockTime int32 `json:"block_time"`

	// 规则描述
	Description *string `json:"description,omitempty"`
}

func (o CreatePunishmentRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePunishmentRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePunishmentRuleRequestBody", string(data)}, " ")
}

type CreatePunishmentRuleRequestBodyCategory struct {
	value string
}

type CreatePunishmentRuleRequestBodyCategoryEnum struct {
	LONG_IP_BLOCK      CreatePunishmentRuleRequestBodyCategory
	LONG_COOKIE_BLOCK  CreatePunishmentRuleRequestBodyCategory
	LONG_PARAMS_BLOCK  CreatePunishmentRuleRequestBodyCategory
	SHORT_IP_BLOCK     CreatePunishmentRuleRequestBodyCategory
	SHORT_COOKIE_BLOCK CreatePunishmentRuleRequestBodyCategory
	SHORT_PARAMS_BLOCK CreatePunishmentRuleRequestBodyCategory
}

func GetCreatePunishmentRuleRequestBodyCategoryEnum() CreatePunishmentRuleRequestBodyCategoryEnum {
	return CreatePunishmentRuleRequestBodyCategoryEnum{
		LONG_IP_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "long_ip_block",
		},
		LONG_COOKIE_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "long_cookie_block",
		},
		LONG_PARAMS_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "long_params_block",
		},
		SHORT_IP_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "short_ip_block",
		},
		SHORT_COOKIE_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "short_cookie_block",
		},
		SHORT_PARAMS_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "short_params_block",
		},
	}
}

func (c CreatePunishmentRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c CreatePunishmentRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePunishmentRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
