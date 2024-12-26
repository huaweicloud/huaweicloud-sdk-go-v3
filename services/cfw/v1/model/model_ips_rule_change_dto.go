package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type IpsRuleChangeDto struct {

	// ips的id列表
	IpsIds *[]string `json:"ips_ids,omitempty"`

	// 防护对象id
	ObjectId *string `json:"object_id,omitempty"`

	// ips规则状态
	Status *IpsRuleChangeDtoStatus `json:"status,omitempty"`
}

func (o IpsRuleChangeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsRuleChangeDto struct{}"
	}

	return strings.Join([]string{"IpsRuleChangeDto", string(data)}, " ")
}

type IpsRuleChangeDtoStatus struct {
	value string
}

type IpsRuleChangeDtoStatusEnum struct {
	OBSERVE     IpsRuleChangeDtoStatus
	ENABLE      IpsRuleChangeDtoStatus
	CLOSE       IpsRuleChangeDtoStatus
	DEFAULT     IpsRuleChangeDtoStatus
	ALL_DEFAULT IpsRuleChangeDtoStatus
}

func GetIpsRuleChangeDtoStatusEnum() IpsRuleChangeDtoStatusEnum {
	return IpsRuleChangeDtoStatusEnum{
		OBSERVE: IpsRuleChangeDtoStatus{
			value: "OBSERVE",
		},
		ENABLE: IpsRuleChangeDtoStatus{
			value: "ENABLE",
		},
		CLOSE: IpsRuleChangeDtoStatus{
			value: "CLOSE",
		},
		DEFAULT: IpsRuleChangeDtoStatus{
			value: "DEFAULT",
		},
		ALL_DEFAULT: IpsRuleChangeDtoStatus{
			value: "ALL_DEFAULT",
		},
	}
}

func (c IpsRuleChangeDtoStatus) Value() string {
	return c.value
}

func (c IpsRuleChangeDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IpsRuleChangeDtoStatus) UnmarshalJSON(b []byte) error {
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
