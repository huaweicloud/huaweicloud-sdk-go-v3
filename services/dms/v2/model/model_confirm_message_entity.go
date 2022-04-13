package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ConfirmMessageEntity struct {
	// 消费时返回的ID。

	Handler *string `json:"handler,omitempty"`
	// 客户端处理数据的状态。 取值为“success”或者“fail”。

	Status *ConfirmMessageEntityStatus `json:"status,omitempty"`
}

func (o ConfirmMessageEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmMessageEntity struct{}"
	}

	return strings.Join([]string{"ConfirmMessageEntity", string(data)}, " ")
}

type ConfirmMessageEntityStatus struct {
	value string
}

type ConfirmMessageEntityStatusEnum struct {
	SUCCESS ConfirmMessageEntityStatus
	FAIL    ConfirmMessageEntityStatus
}

func GetConfirmMessageEntityStatusEnum() ConfirmMessageEntityStatusEnum {
	return ConfirmMessageEntityStatusEnum{
		SUCCESS: ConfirmMessageEntityStatus{
			value: "success",
		},
		FAIL: ConfirmMessageEntityStatus{
			value: "fail",
		},
	}
}

func (c ConfirmMessageEntityStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmMessageEntityStatus) UnmarshalJSON(b []byte) error {
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
