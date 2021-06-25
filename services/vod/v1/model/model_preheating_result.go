package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PreheatingResult struct {
	// 预热任务数组。<br/>

	Url *string `json:"url,omitempty"`
	// 预热任务状态，取值processing， succeed， failed，分别表示处理中，完成，失败。<br/>

	Status *PreheatingResultStatus `json:"status,omitempty"`
}

func (o PreheatingResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PreheatingResult struct{}"
	}

	return strings.Join([]string{"PreheatingResult", string(data)}, " ")
}

type PreheatingResultStatus struct {
	value string
}

type PreheatingResultStatusEnum struct {
	PROCESSING PreheatingResultStatus
	SUCCEED    PreheatingResultStatus
	FAILED     PreheatingResultStatus
}

func GetPreheatingResultStatusEnum() PreheatingResultStatusEnum {
	return PreheatingResultStatusEnum{
		PROCESSING: PreheatingResultStatus{
			value: "PROCESSING",
		},
		SUCCEED: PreheatingResultStatus{
			value: "SUCCEED",
		},
		FAILED: PreheatingResultStatus{
			value: "FAILED",
		},
	}
}

func (c PreheatingResultStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *PreheatingResultStatus) UnmarshalJSON(b []byte) error {
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
