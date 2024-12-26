package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SinkObsParameters struct {

	// AK
	AccessKey string `json:"access_key"`

	// SK
	SecretKey string `json:"secret_key"`

	// 桶
	ObsBucket string `json:"obs_bucket"`

	// 转储目录
	ObsPath *string `json:"obs_path,omitempty"`

	// 时间目录格式
	TimeFormat SinkObsParametersTimeFormat `json:"time_format"`
}

func (o SinkObsParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SinkObsParameters struct{}"
	}

	return strings.Join([]string{"SinkObsParameters", string(data)}, " ")
}

type SinkObsParametersTimeFormat struct {
	value string
}

type SinkObsParametersTimeFormatEnum struct {
	YYYY             SinkObsParametersTimeFormat
	YYYY_MM          SinkObsParametersTimeFormat
	YYYY_MM_DD       SinkObsParametersTimeFormat
	YYYY_MM_DD_HH    SinkObsParametersTimeFormat
	YYYY_MM_DD_HH_MM SinkObsParametersTimeFormat
}

func GetSinkObsParametersTimeFormatEnum() SinkObsParametersTimeFormatEnum {
	return SinkObsParametersTimeFormatEnum{
		YYYY: SinkObsParametersTimeFormat{
			value: "yyyy",
		},
		YYYY_MM: SinkObsParametersTimeFormat{
			value: "yyyy/MM",
		},
		YYYY_MM_DD: SinkObsParametersTimeFormat{
			value: "yyyy/MM/dd",
		},
		YYYY_MM_DD_HH: SinkObsParametersTimeFormat{
			value: "yyyy/MM/dd/HH",
		},
		YYYY_MM_DD_HH_MM: SinkObsParametersTimeFormat{
			value: "yyyy/MM/dd/HH/mm",
		},
	}
}

func (c SinkObsParametersTimeFormat) Value() string {
	return c.value
}

func (c SinkObsParametersTimeFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SinkObsParametersTimeFormat) UnmarshalJSON(b []byte) error {
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
