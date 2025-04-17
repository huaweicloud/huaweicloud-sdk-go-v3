package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DailyLog EIP异常事件响应体
type DailyLog struct {

	// 开始时间
	StartTime int64 `json:"start_time"`

	// 结束时间
	EndTime int64 `json:"end_time"`

	// 防护状态，可选范围： - 1：表示清洗 - 2：表示黑洞
	Status DailyLogStatus `json:"status"`

	// 触发时流量
	TriggerBps int32 `json:"trigger_bps"`

	// 触发时报文速率
	TriggerPps int32 `json:"trigger_pps"`

	// 触发时HTTP请求速率
	TriggerHttpPps int32 `json:"trigger_http_pps"`
}

func (o DailyLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DailyLog struct{}"
	}

	return strings.Join([]string{"DailyLog", string(data)}, " ")
}

type DailyLogStatus struct {
	value int32
}

type DailyLogStatusEnum struct {
	E_1 DailyLogStatus
	E_2 DailyLogStatus
}

func GetDailyLogStatusEnum() DailyLogStatusEnum {
	return DailyLogStatusEnum{
		E_1: DailyLogStatus{
			value: 1,
		}, E_2: DailyLogStatus{
			value: 2,
		},
	}
}

func (c DailyLogStatus) Value() int32 {
	return c.value
}

func (c DailyLogStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DailyLogStatus) UnmarshalJSON(b []byte) error {
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
