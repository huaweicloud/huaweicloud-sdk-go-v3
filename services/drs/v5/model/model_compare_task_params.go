package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CompareTaskParams 创建或取消对比任务信息体。
type CompareTaskParams struct {

	// 取消对比任务必填。
	CompareTaskId *string `json:"compare_task_id,omitempty"`

	// 对比任务模式。取值： - object：对象对比。 - lines：行数对比。 - contents：内容对比。
	Type *CompareTaskParamsType `json:"type,omitempty"`

	// 定时启动时间，时间戳格式。
	StartTime *string `json:"start_time,omitempty"`

	// 对比策略。
	Option map[string]string `json:"option,omitempty"`

	// 对比选择对象。
	DbObject map[string]DatabaseObject `json:"db_object,omitempty"`

	// 更新数据加工规则请求体
	DataProcessInfo *[]DataProcessInfo `json:"data_process_info,omitempty"`
}

func (o CompareTaskParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareTaskParams struct{}"
	}

	return strings.Join([]string{"CompareTaskParams", string(data)}, " ")
}

type CompareTaskParamsType struct {
	value string
}

type CompareTaskParamsTypeEnum struct {
	OBJECT   CompareTaskParamsType
	LINES    CompareTaskParamsType
	CONTENTS CompareTaskParamsType
}

func GetCompareTaskParamsTypeEnum() CompareTaskParamsTypeEnum {
	return CompareTaskParamsTypeEnum{
		OBJECT: CompareTaskParamsType{
			value: "object",
		},
		LINES: CompareTaskParamsType{
			value: "lines",
		},
		CONTENTS: CompareTaskParamsType{
			value: "contents",
		},
	}
}

func (c CompareTaskParamsType) Value() string {
	return c.value
}

func (c CompareTaskParamsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CompareTaskParamsType) UnmarshalJSON(b []byte) error {
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
