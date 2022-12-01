package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddRuleAclDto
type AddRuleAclDto struct {

	// 防护对象id
	ObjectId string `json:"object_id"`

	// 规则type，0：互联网规则，1:vpc规则，2：nat规则
	Type AddRuleAclDtoType `json:"type"`

	// rules
	Rules []AddRuleAclDtoRules `json:"rules"`
}

func (o AddRuleAclDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddRuleAclDto struct{}"
	}

	return strings.Join([]string{"AddRuleAclDto", string(data)}, " ")
}

type AddRuleAclDtoType struct {
	value int32
}

type AddRuleAclDtoTypeEnum struct {
	E_0 AddRuleAclDtoType
	E_1 AddRuleAclDtoType
	E_2 AddRuleAclDtoType
}

func GetAddRuleAclDtoTypeEnum() AddRuleAclDtoTypeEnum {
	return AddRuleAclDtoTypeEnum{
		E_0: AddRuleAclDtoType{
			value: 0,
		}, E_1: AddRuleAclDtoType{
			value: 1,
		}, E_2: AddRuleAclDtoType{
			value: 2,
		},
	}
}

func (c AddRuleAclDtoType) Value() int32 {
	return c.value
}

func (c AddRuleAclDtoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddRuleAclDtoType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
