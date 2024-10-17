package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchSwitchesRequest struct {

	// 规则ID,多个ID中间逗号分隔。可在查询风险规则策略接口ID字段获取。
	Ids *string `json:"ids,omitempty"`

	// 开关状态 - OFF: 关闭 - ON: 开启
	Status *BatchSwitchesRequestStatus `json:"status,omitempty"`
}

func (o BatchSwitchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSwitchesRequest struct{}"
	}

	return strings.Join([]string{"BatchSwitchesRequest", string(data)}, " ")
}

type BatchSwitchesRequestStatus struct {
	value string
}

type BatchSwitchesRequestStatusEnum struct {
	OFF BatchSwitchesRequestStatus
	ON  BatchSwitchesRequestStatus
}

func GetBatchSwitchesRequestStatusEnum() BatchSwitchesRequestStatusEnum {
	return BatchSwitchesRequestStatusEnum{
		OFF: BatchSwitchesRequestStatus{
			value: "OFF",
		},
		ON: BatchSwitchesRequestStatus{
			value: "ON",
		},
	}
}

func (c BatchSwitchesRequestStatus) Value() string {
	return c.value
}

func (c BatchSwitchesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSwitchesRequestStatus) UnmarshalJSON(b []byte) error {
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
