package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 组件创建信息
type Detail struct {
	// 开始时间

	StartTime *string `json:"start_time,omitempty"`
	// 结束时间

	EndTime *string `json:"end_time,omitempty"`
	// 状态

	Status *DetailStatus `json:"status,omitempty"`
	// 细节描述

	Detail *string `json:"detail,omitempty"`
}

func (o Detail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Detail struct{}"
	}

	return strings.Join([]string{"Detail", string(data)}, " ")
}

type DetailStatus struct {
	value string
}

type DetailStatusEnum struct {
	WAITING   DetailStatus
	DEPLOYING DetailStatus
	FINISHED  DetailStatus
	FAILED    DetailStatus
}

func GetDetailStatusEnum() DetailStatusEnum {
	return DetailStatusEnum{
		WAITING: DetailStatus{
			value: "waiting",
		},
		DEPLOYING: DetailStatus{
			value: "deploying",
		},
		FINISHED: DetailStatus{
			value: "finished",
		},
		FAILED: DetailStatus{
			value: "failed",
		},
	}
}

func (c DetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DetailStatus) UnmarshalJSON(b []byte) error {
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
