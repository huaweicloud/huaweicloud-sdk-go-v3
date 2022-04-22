package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchListModulesRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数，默认值为10，取值区间为1-1000
	Limit *int32 `json:"limit,omitempty"`

	// 应用类型
	AppType *BatchListModulesRequestAppType `json:"app_type,omitempty"`

	// 功能类型
	FunctionType *BatchListModulesRequestFunctionType `json:"function_type,omitempty"`
}

func (o BatchListModulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListModulesRequest struct{}"
	}

	return strings.Join([]string{"BatchListModulesRequest", string(data)}, " ")
}

type BatchListModulesRequestAppType struct {
	value string
}

type BatchListModulesRequestAppTypeEnum struct {
	SYSTEM_REQUIRED BatchListModulesRequestAppType
	SYSTEM_OPTIONAL BatchListModulesRequestAppType
	USER            BatchListModulesRequestAppType
}

func GetBatchListModulesRequestAppTypeEnum() BatchListModulesRequestAppTypeEnum {
	return BatchListModulesRequestAppTypeEnum{
		SYSTEM_REQUIRED: BatchListModulesRequestAppType{
			value: "SYSTEM_REQUIRED",
		},
		SYSTEM_OPTIONAL: BatchListModulesRequestAppType{
			value: "SYSTEM_OPTIONAL",
		},
		USER: BatchListModulesRequestAppType{
			value: "USER",
		},
	}
}

func (c BatchListModulesRequestAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListModulesRequestAppType) UnmarshalJSON(b []byte) error {
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

type BatchListModulesRequestFunctionType struct {
	value string
}

type BatchListModulesRequestFunctionTypeEnum struct {
	DATA_PROCESSING        BatchListModulesRequestFunctionType
	PROTOCOL_PARSING       BatchListModulesRequestFunctionType
	ON_PREMISE_INTEGRATION BatchListModulesRequestFunctionType
}

func GetBatchListModulesRequestFunctionTypeEnum() BatchListModulesRequestFunctionTypeEnum {
	return BatchListModulesRequestFunctionTypeEnum{
		DATA_PROCESSING: BatchListModulesRequestFunctionType{
			value: "DATA_PROCESSING",
		},
		PROTOCOL_PARSING: BatchListModulesRequestFunctionType{
			value: "PROTOCOL_PARSING",
		},
		ON_PREMISE_INTEGRATION: BatchListModulesRequestFunctionType{
			value: "ON_PREMISE_INTEGRATION",
		},
	}
}

func (c BatchListModulesRequestFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListModulesRequestFunctionType) UnmarshalJSON(b []byte) error {
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
