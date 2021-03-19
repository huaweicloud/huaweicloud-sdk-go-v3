package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BlackEnhance struct {
	// 提前反馈“疑似黑边”开关，取值为on或off。

	EarlyReport *BlackEnhanceEarlyReport `json:"early_report,omitempty"`
	// 参数格式：top:bottom:left:right

	Position *string `json:"position,omitempty"`
	// 黑边剪裁检测起始时间

	StartTime *string `json:"start_time,omitempty"`
}

func (o BlackEnhance) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BlackEnhance struct{}"
	}

	return strings.Join([]string{"BlackEnhance", string(data)}, " ")
}

type BlackEnhanceEarlyReport struct {
	value string
}

type BlackEnhanceEarlyReportEnum struct {
	ON  BlackEnhanceEarlyReport
	OFF BlackEnhanceEarlyReport
}

func GetBlackEnhanceEarlyReportEnum() BlackEnhanceEarlyReportEnum {
	return BlackEnhanceEarlyReportEnum{
		ON: BlackEnhanceEarlyReport{
			value: "on",
		},
		OFF: BlackEnhanceEarlyReport{
			value: "off",
		},
	}
}

func (c BlackEnhanceEarlyReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BlackEnhanceEarlyReport) UnmarshalJSON(b []byte) error {
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
