package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type PolicyTriggerResp struct {
	// 调度器id

	Id string `json:"id"`
	// 调度器名称

	Name *string `json:"name,omitempty"`

	Properties *PolicyTriggerPropertiesResp `json:"properties"`
	// 调度器类型,目前只支持 time: 定时调度。

	Type *PolicyTriggerRespType `json:"type,omitempty"`
}

func (o PolicyTriggerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyTriggerResp struct{}"
	}

	return strings.Join([]string{"PolicyTriggerResp", string(data)}, " ")
}

type PolicyTriggerRespType struct {
	value string
}

type PolicyTriggerRespTypeEnum struct {
	TIME PolicyTriggerRespType
}

func GetPolicyTriggerRespTypeEnum() PolicyTriggerRespTypeEnum {
	return PolicyTriggerRespTypeEnum{
		TIME: PolicyTriggerRespType{
			value: "time",
		},
	}
}

func (c PolicyTriggerRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyTriggerRespType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
