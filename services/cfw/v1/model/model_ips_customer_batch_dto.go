package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IpsCustomerBatchDto **参数解释：** 批量自定义IPS规则请求体 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type IpsCustomerBatchDto struct {

	// **参数解释**： 自定义IPS规则执行动作,仅更新自定义IPS规则场景下需要设置，其他场景无需设置 **约束限制**：   不涉及 **取值范围**： 0：只记录日志 1：重置/拦截 **默认取值**：   不涉及
	ActionType *IpsCustomerBatchDtoActionType `json:"action_type,omitempty"`

	// **参数解释**： ips规则ID， 可通过调用获取ips规则列表获取，通过data.records.ips_id获取。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	IpsIds []string `json:"ips_ids"`

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId string `json:"fw_instance_id"`
}

func (o IpsCustomerBatchDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsCustomerBatchDto struct{}"
	}

	return strings.Join([]string{"IpsCustomerBatchDto", string(data)}, " ")
}

type IpsCustomerBatchDtoActionType struct {
	value int32
}

type IpsCustomerBatchDtoActionTypeEnum struct {
	E_0 IpsCustomerBatchDtoActionType
	E_1 IpsCustomerBatchDtoActionType
}

func GetIpsCustomerBatchDtoActionTypeEnum() IpsCustomerBatchDtoActionTypeEnum {
	return IpsCustomerBatchDtoActionTypeEnum{
		E_0: IpsCustomerBatchDtoActionType{
			value: 0,
		}, E_1: IpsCustomerBatchDtoActionType{
			value: 1,
		},
	}
}

func (c IpsCustomerBatchDtoActionType) Value() int32 {
	return c.value
}

func (c IpsCustomerBatchDtoActionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsCustomerBatchDtoActionType) UnmarshalJSON(b []byte) error {
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
