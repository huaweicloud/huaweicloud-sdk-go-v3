package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ChangeProtectStatusRequestBody struct {

	// **参数解释**： 防护对象ID，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)获得，通过返回值中的data.records.protect_objects.object_id（.表示各对象之间层级的区分）获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	ObjectId string `json:"object_id"`

	// **参数解释**： VPC边界防护状态 **约束限制**： 不涉及 **取值范围**： 0- 开启防护、1- 关闭防护 **默认取值**： 不涉及
	Status ChangeProtectStatusRequestBodyStatus `json:"status"`
}

func (o ChangeProtectStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeProtectStatusRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeProtectStatusRequestBody", string(data)}, " ")
}

type ChangeProtectStatusRequestBodyStatus struct {
	value int32
}

type ChangeProtectStatusRequestBodyStatusEnum struct {
	E_0 ChangeProtectStatusRequestBodyStatus
	E_1 ChangeProtectStatusRequestBodyStatus
}

func GetChangeProtectStatusRequestBodyStatusEnum() ChangeProtectStatusRequestBodyStatusEnum {
	return ChangeProtectStatusRequestBodyStatusEnum{
		E_0: ChangeProtectStatusRequestBodyStatus{
			value: 0,
		}, E_1: ChangeProtectStatusRequestBodyStatus{
			value: 1,
		},
	}
}

func (c ChangeProtectStatusRequestBodyStatus) Value() int32 {
	return c.value
}

func (c ChangeProtectStatusRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeProtectStatusRequestBodyStatus) UnmarshalJSON(b []byte) error {
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
