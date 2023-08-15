package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdatePunishmentRuleRequestBody struct {

	// 攻击惩罚类别
	Category UpdatePunishmentRuleRequestBodyCategory `json:"category"`

	// 拦截时间，如果选择前缀为long的攻击惩罚类别，则block_time时长范围设置为301-1800;选择前缀为short的攻击惩罚类别，则block_time时长范围为0-300之间
	BlockTime int32 `json:"block_time"`

	// 规则描述
	Description *string `json:"description,omitempty"`
}

func (o UpdatePunishmentRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePunishmentRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePunishmentRuleRequestBody", string(data)}, " ")
}

type UpdatePunishmentRuleRequestBodyCategory struct {
	value string
}

type UpdatePunishmentRuleRequestBodyCategoryEnum struct {
	LONG_IP_BLOCK      UpdatePunishmentRuleRequestBodyCategory
	LONG_COOKIE_BLOCK  UpdatePunishmentRuleRequestBodyCategory
	LONG_PARAMS_BLOCK  UpdatePunishmentRuleRequestBodyCategory
	SHORT_IP_BLOCK     UpdatePunishmentRuleRequestBodyCategory
	SHORT_COOKIE_BLOCK UpdatePunishmentRuleRequestBodyCategory
	SHORT_PARAMS_BLOCK UpdatePunishmentRuleRequestBodyCategory
}

func GetUpdatePunishmentRuleRequestBodyCategoryEnum() UpdatePunishmentRuleRequestBodyCategoryEnum {
	return UpdatePunishmentRuleRequestBodyCategoryEnum{
		LONG_IP_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "long_ip_block",
		},
		LONG_COOKIE_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "long_cookie_block",
		},
		LONG_PARAMS_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "long_params_block",
		},
		SHORT_IP_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "short_ip_block",
		},
		SHORT_COOKIE_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "short_cookie_block",
		},
		SHORT_PARAMS_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "short_params_block",
		},
	}
}

func (c UpdatePunishmentRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c UpdatePunishmentRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePunishmentRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
