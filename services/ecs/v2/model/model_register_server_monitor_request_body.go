package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto create Body Object
type RegisterServerMonitorRequestBody struct {

	// 注册云服务器监控。
	MonitorMetrics RegisterServerMonitorRequestBodyMonitorMetrics `json:"monitorMetrics"`
}

func (o RegisterServerMonitorRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerMonitorRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterServerMonitorRequestBody", string(data)}, " ")
}

type RegisterServerMonitorRequestBodyMonitorMetrics struct {
	value string
}

type RegisterServerMonitorRequestBodyMonitorMetricsEnum struct {
	EMPTY RegisterServerMonitorRequestBodyMonitorMetrics
}

func GetRegisterServerMonitorRequestBodyMonitorMetricsEnum() RegisterServerMonitorRequestBodyMonitorMetricsEnum {
	return RegisterServerMonitorRequestBodyMonitorMetricsEnum{
		EMPTY: RegisterServerMonitorRequestBodyMonitorMetrics{
			value: "",
		},
	}
}

func (c RegisterServerMonitorRequestBodyMonitorMetrics) Value() string {
	return c.value
}

func (c RegisterServerMonitorRequestBodyMonitorMetrics) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RegisterServerMonitorRequestBodyMonitorMetrics) UnmarshalJSON(b []byte) error {
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
