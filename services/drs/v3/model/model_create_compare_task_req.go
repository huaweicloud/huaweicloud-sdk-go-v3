package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type CreateCompareTaskReq struct {

	// 任务id。
	JobId string `json:"job_id" xml:"job_id"`

	// 对象级对比类型，取值为空代表不创建对象级对比。object_level_compare_type和data_level_compare_info都为空时，只查询已创建的对比任务列表。
	ObjectLevelCompareType *CreateCompareTaskReqObjectLevelCompareType `json:"object_level_compare_type,omitempty" xml:"object_level_compare_type"`

	DataLevelCompareInfo *CreateDataLevelCompareReq `json:"data_level_compare_info,omitempty" xml:"data_level_compare_info"`
}

func (o CreateCompareTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompareTaskReq struct{}"
	}

	return strings.Join([]string{"CreateCompareTaskReq", string(data)}, " ")
}

type CreateCompareTaskReqObjectLevelCompareType struct {
	value string
}

type CreateCompareTaskReqObjectLevelCompareTypeEnum struct {
	OBJECTS CreateCompareTaskReqObjectLevelCompareType
}

func GetCreateCompareTaskReqObjectLevelCompareTypeEnum() CreateCompareTaskReqObjectLevelCompareTypeEnum {
	return CreateCompareTaskReqObjectLevelCompareTypeEnum{
		OBJECTS: CreateCompareTaskReqObjectLevelCompareType{
			value: "objects",
		},
	}
}

func (c CreateCompareTaskReqObjectLevelCompareType) Value() string {
	return c.value
}

func (c CreateCompareTaskReqObjectLevelCompareType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCompareTaskReqObjectLevelCompareType) UnmarshalJSON(b []byte) error {
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
