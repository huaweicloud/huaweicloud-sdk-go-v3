package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchSwitchesRequestNew struct {

	// 规则ID列表
	Ids []string `json:"ids"`

	// 开关状态 - OFF：关闭 - ON：开启
	Status BatchSwitchesRequestNewStatus `json:"status"`
}

func (o BatchSwitchesRequestNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSwitchesRequestNew struct{}"
	}

	return strings.Join([]string{"BatchSwitchesRequestNew", string(data)}, " ")
}

type BatchSwitchesRequestNewStatus struct {
	value string
}

type BatchSwitchesRequestNewStatusEnum struct {
	OFF BatchSwitchesRequestNewStatus
	ON  BatchSwitchesRequestNewStatus
}

func GetBatchSwitchesRequestNewStatusEnum() BatchSwitchesRequestNewStatusEnum {
	return BatchSwitchesRequestNewStatusEnum{
		OFF: BatchSwitchesRequestNewStatus{
			value: "OFF",
		},
		ON: BatchSwitchesRequestNewStatus{
			value: "ON",
		},
	}
}

func (c BatchSwitchesRequestNewStatus) Value() string {
	return c.value
}

func (c BatchSwitchesRequestNewStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSwitchesRequestNewStatus) UnmarshalJSON(b []byte) error {
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
