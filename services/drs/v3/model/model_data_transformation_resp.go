package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataTransformationResp 数据加工返回体
type DataTransformationResp struct {

	// 任务id
	Id *string `json:"id,omitempty"`

	// 状态
	Status *DataTransformationRespStatus `json:"status,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o DataTransformationResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataTransformationResp struct{}"
	}

	return strings.Join([]string{"DataTransformationResp", string(data)}, " ")
}

type DataTransformationRespStatus struct {
	value string
}

type DataTransformationRespStatusEnum struct {
	SUCCESS DataTransformationRespStatus
	FAILED  DataTransformationRespStatus
}

func GetDataTransformationRespStatusEnum() DataTransformationRespStatusEnum {
	return DataTransformationRespStatusEnum{
		SUCCESS: DataTransformationRespStatus{
			value: "success",
		},
		FAILED: DataTransformationRespStatus{
			value: "failed",
		},
	}
}

func (c DataTransformationRespStatus) Value() string {
	return c.value
}

func (c DataTransformationRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataTransformationRespStatus) UnmarshalJSON(b []byte) error {
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
