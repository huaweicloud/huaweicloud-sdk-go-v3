package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobOutputResp 作业输出。
type JobOutputResp struct {

	// **参数解释**：输出数据的名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：输出项类型。 **取值范围**：枚举值如下： - obs：OBS - model：AI应用元模型
	Type *JobOutputRespType `json:"type,omitempty"`

	// **参数解释**：输出配置。
	Config map[string]interface{} `json:"config,omitempty"`
}

func (o JobOutputResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobOutputResp struct{}"
	}

	return strings.Join([]string{"JobOutputResp", string(data)}, " ")
}

type JobOutputRespType struct {
	value string
}

type JobOutputRespTypeEnum struct {
	DATASET JobOutputRespType
	OBSOBS  JobOutputRespType
}

func GetJobOutputRespTypeEnum() JobOutputRespTypeEnum {
	return JobOutputRespTypeEnum{
		DATASET: JobOutputRespType{
			value: "dataset：数据集。",
		},
		OBSOBS: JobOutputRespType{
			value: "obs：OBS文件",
		},
	}
}

func (c JobOutputRespType) Value() string {
	return c.value
}

func (c JobOutputRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobOutputRespType) UnmarshalJSON(b []byte) error {
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
