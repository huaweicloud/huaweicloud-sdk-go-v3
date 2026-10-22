package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateBotMRuleActionRequestBody struct {

	// 防护动作对应的数字
	DefenseAction *UpdateBotMRuleActionRequestBodyDefenseAction `json:"defense_action,omitempty"`
}

func (o UpdateBotMRuleActionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBotMRuleActionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateBotMRuleActionRequestBody", string(data)}, " ")
}

type UpdateBotMRuleActionRequestBodyDefenseAction struct {
	value int32
}

type UpdateBotMRuleActionRequestBodyDefenseActionEnum struct {
	E_101 UpdateBotMRuleActionRequestBodyDefenseAction
	E_102 UpdateBotMRuleActionRequestBodyDefenseAction
	E_103 UpdateBotMRuleActionRequestBodyDefenseAction
	E_104 UpdateBotMRuleActionRequestBodyDefenseAction
	E_201 UpdateBotMRuleActionRequestBodyDefenseAction
	E_202 UpdateBotMRuleActionRequestBodyDefenseAction
	E_301 UpdateBotMRuleActionRequestBodyDefenseAction
}

func GetUpdateBotMRuleActionRequestBodyDefenseActionEnum() UpdateBotMRuleActionRequestBodyDefenseActionEnum {
	return UpdateBotMRuleActionRequestBodyDefenseActionEnum{
		E_101: UpdateBotMRuleActionRequestBodyDefenseAction{
			value: 101,
		}, E_102: UpdateBotMRuleActionRequestBodyDefenseAction{
			value: 102,
		}, E_103: UpdateBotMRuleActionRequestBodyDefenseAction{
			value: 103,
		}, E_104: UpdateBotMRuleActionRequestBodyDefenseAction{
			value: 104,
		}, E_201: UpdateBotMRuleActionRequestBodyDefenseAction{
			value: 201,
		}, E_202: UpdateBotMRuleActionRequestBodyDefenseAction{
			value: 202,
		}, E_301: UpdateBotMRuleActionRequestBodyDefenseAction{
			value: 301,
		},
	}
}

func (c UpdateBotMRuleActionRequestBodyDefenseAction) Value() int32 {
	return c.value
}

func (c UpdateBotMRuleActionRequestBodyDefenseAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateBotMRuleActionRequestBodyDefenseAction) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
