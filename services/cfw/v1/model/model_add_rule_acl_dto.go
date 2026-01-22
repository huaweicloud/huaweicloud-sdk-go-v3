package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddRuleAclDto AddRuleAclDto
type AddRuleAclDto struct {

	// **参数解释**： 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志ID，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得 **约束限制**： type为0时，object_id为互联网边界防护对象ID，type为1时，object_id为VPC边界防护对象ID，type可通过data.records.protect_objects.type（.表示各对象之间层级的区分）获得 **取值范围**： 32位UUID **默认取值**： 不涉及
	ObjectId string `json:"object_id"`

	// **参数解释**： 规则类型，用于区分不同防护对象设置规则类型。 **约束限制**： 不涉及 **取值范围**：  0：互联网边界规则，源（source）和目的（destination）地址需要为公网IP或域名； 1：VPC间规则，源（source）和目的（destination）地址需要为私有ip； 2：NAT规则，源（source）地址需要为私网IP，目的地址为公网IP或域名。 **默认取值**： 不涉及
	Type AddRuleAclDtoType `json:"type"`

	// **参数解释**： 添加规则请求规则列表 **约束限制**： 不涉及
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
