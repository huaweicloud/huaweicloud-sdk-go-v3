package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StatisticsSourceDto **参数说明**：消息来源。
type StatisticsSourceDto struct {

	// **参数说明**：信息来源的具体类型描述。
	SourceType *StatisticsSourceDtoSourceType `json:"source_type,omitempty"`

	// **参数说明**：信息来源的唯一标识码ID。
	SourceId *string `json:"source_id,omitempty"`
}

func (o StatisticsSourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsSourceDto struct{}"
	}

	return strings.Join([]string{"StatisticsSourceDto", string(data)}, " ")
}

type StatisticsSourceDtoSourceType struct {
	value string
}

type StatisticsSourceDtoSourceTypeEnum struct {
	RSU       StatisticsSourceDtoSourceType
	OBU       StatisticsSourceDtoSourceType
	DETECTION StatisticsSourceDtoSourceType
}

func GetStatisticsSourceDtoSourceTypeEnum() StatisticsSourceDtoSourceTypeEnum {
	return StatisticsSourceDtoSourceTypeEnum{
		RSU: StatisticsSourceDtoSourceType{
			value: "rsu",
		},
		OBU: StatisticsSourceDtoSourceType{
			value: "obu",
		},
		DETECTION: StatisticsSourceDtoSourceType{
			value: "detection",
		},
	}
}

func (c StatisticsSourceDtoSourceType) Value() string {
	return c.value
}

func (c StatisticsSourceDtoSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatisticsSourceDtoSourceType) UnmarshalJSON(b []byte) error {
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
