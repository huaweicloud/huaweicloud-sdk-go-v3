package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobOutput 作业输出。
type JobOutput struct {

	// 输出数据的名称。
	Name *string `json:"name,omitempty"`

	// 输出项类型。枚举值如下： - obs：OBS - model：AI应用元模型
	Type *JobOutputType `json:"type,omitempty"`

	// 输出配置。
	Config map[string]interface{} `json:"config,omitempty"`
}

func (o JobOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobOutput struct{}"
	}

	return strings.Join([]string{"JobOutput", string(data)}, " ")
}

type JobOutputType struct {
	value string
}

type JobOutputTypeEnum struct {
	DATASET JobOutputType
	OBSOBS  JobOutputType
}

func GetJobOutputTypeEnum() JobOutputTypeEnum {
	return JobOutputTypeEnum{
		DATASET: JobOutputType{
			value: "dataset：数据集。",
		},
		OBSOBS: JobOutputType{
			value: "obs：OBS文件",
		},
	}
}

func (c JobOutputType) Value() string {
	return c.value
}

func (c JobOutputType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobOutputType) UnmarshalJSON(b []byte) error {
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
