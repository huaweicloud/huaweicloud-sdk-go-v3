package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddRuleAclDto AddRuleAclDto
type AddRuleAclDto struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得
	ObjectId string `json:"object_id"`

	// 规则类型，0：互联网边界规则，1：vpc间规则，2：nat规则，当type取0时，规则源和目的地址需要为公网ip或域名，vpc间规则需要源和目的地址为私有ip，nat规则需要源地址为私网ip，目的地址为公网ip或域名。
	Type AddRuleAclDtoType `json:"type"`

	// 添加规则请求规则列表
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
