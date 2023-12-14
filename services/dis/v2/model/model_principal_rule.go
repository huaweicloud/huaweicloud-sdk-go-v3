package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PrincipalRule struct {

	// 授权用户ID。
	Principal *string `json:"principal,omitempty"`

	// 授权用户名。  如果授权给租户下的所有子用户，格式为：domainName.*；如果授权给租户下的指定子用户，则格式为：domainName.userName
	PrincipalName *string `json:"principal_name,omitempty"`

	// 授权操作类型。  - putRecords：上传数据。 - getRecords：下载数据。
	ActionType *PrincipalRuleActionType `json:"action_type,omitempty"`

	// 授权影响类型。  - accept：允许该授权操作。
	Effect *PrincipalRuleEffect `json:"effect,omitempty"`
}

func (o PrincipalRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrincipalRule struct{}"
	}

	return strings.Join([]string{"PrincipalRule", string(data)}, " ")
}

type PrincipalRuleActionType struct {
	value string
}

type PrincipalRuleActionTypeEnum struct {
	PUT_RECORDS PrincipalRuleActionType
	GET_RECORDS PrincipalRuleActionType
}

func GetPrincipalRuleActionTypeEnum() PrincipalRuleActionTypeEnum {
	return PrincipalRuleActionTypeEnum{
		PUT_RECORDS: PrincipalRuleActionType{
			value: "putRecords",
		},
		GET_RECORDS: PrincipalRuleActionType{
			value: "getRecords",
		},
	}
}

func (c PrincipalRuleActionType) Value() string {
	return c.value
}

func (c PrincipalRuleActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrincipalRuleActionType) UnmarshalJSON(b []byte) error {
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

type PrincipalRuleEffect struct {
	value string
}

type PrincipalRuleEffectEnum struct {
	ACCEPT PrincipalRuleEffect
}

func GetPrincipalRuleEffectEnum() PrincipalRuleEffectEnum {
	return PrincipalRuleEffectEnum{
		ACCEPT: PrincipalRuleEffect{
			value: "accept",
		},
	}
}

func (c PrincipalRuleEffect) Value() string {
	return c.value
}

func (c PrincipalRuleEffect) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrincipalRuleEffect) UnmarshalJSON(b []byte) error {
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
