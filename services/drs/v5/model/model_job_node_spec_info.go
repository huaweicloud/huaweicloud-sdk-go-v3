package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 任务实例规格信息体。
type JobNodeSpecInfo struct {

	// 实例规格编码。取值： - micro：极小规格。 - small：小规格。 - medium：中规格。 - high：大规格。
	NodeType JobNodeSpecInfoNodeType `json:"node_type"`
}

func (o JobNodeSpecInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobNodeSpecInfo struct{}"
	}

	return strings.Join([]string{"JobNodeSpecInfo", string(data)}, " ")
}

type JobNodeSpecInfoNodeType struct {
	value string
}

type JobNodeSpecInfoNodeTypeEnum struct {
	MICRO  JobNodeSpecInfoNodeType
	SMALL  JobNodeSpecInfoNodeType
	MEDIUM JobNodeSpecInfoNodeType
	HIGH   JobNodeSpecInfoNodeType
}

func GetJobNodeSpecInfoNodeTypeEnum() JobNodeSpecInfoNodeTypeEnum {
	return JobNodeSpecInfoNodeTypeEnum{
		MICRO: JobNodeSpecInfoNodeType{
			value: "micro",
		},
		SMALL: JobNodeSpecInfoNodeType{
			value: "small",
		},
		MEDIUM: JobNodeSpecInfoNodeType{
			value: "medium",
		},
		HIGH: JobNodeSpecInfoNodeType{
			value: "high",
		},
	}
}

func (c JobNodeSpecInfoNodeType) Value() string {
	return c.value
}

func (c JobNodeSpecInfoNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobNodeSpecInfoNodeType) UnmarshalJSON(b []byte) error {
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
