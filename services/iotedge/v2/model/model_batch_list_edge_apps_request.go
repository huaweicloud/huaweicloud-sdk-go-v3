package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchListEdgeAppsRequest struct {

	// 应用ID搜索关键字
	EdgeAppId *string `json:"edge_app_id,omitempty" xml:"edge_app_id"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页记录数，默认值为10，取值区间为1-1000
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 应用id搜索关键字
	AppType *BatchListEdgeAppsRequestAppType `json:"app_type,omitempty" xml:"app_type"`

	// 功能类型
	FunctionType *BatchListEdgeAppsRequestFunctionType `json:"function_type,omitempty" xml:"function_type"`
}

func (o BatchListEdgeAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListEdgeAppsRequest struct{}"
	}

	return strings.Join([]string{"BatchListEdgeAppsRequest", string(data)}, " ")
}

type BatchListEdgeAppsRequestAppType struct {
	value string
}

type BatchListEdgeAppsRequestAppTypeEnum struct {
	SYSTEM_REQUIRED BatchListEdgeAppsRequestAppType
	SYSTEM_OPTIONAL BatchListEdgeAppsRequestAppType
	USER            BatchListEdgeAppsRequestAppType
}

func GetBatchListEdgeAppsRequestAppTypeEnum() BatchListEdgeAppsRequestAppTypeEnum {
	return BatchListEdgeAppsRequestAppTypeEnum{
		SYSTEM_REQUIRED: BatchListEdgeAppsRequestAppType{
			value: "SYSTEM_REQUIRED",
		},
		SYSTEM_OPTIONAL: BatchListEdgeAppsRequestAppType{
			value: "SYSTEM_OPTIONAL",
		},
		USER: BatchListEdgeAppsRequestAppType{
			value: "USER",
		},
	}
}

func (c BatchListEdgeAppsRequestAppType) Value() string {
	return c.value
}

func (c BatchListEdgeAppsRequestAppType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListEdgeAppsRequestAppType) UnmarshalJSON(b []byte) error {
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

type BatchListEdgeAppsRequestFunctionType struct {
	value string
}

type BatchListEdgeAppsRequestFunctionTypeEnum struct {
	DATA_PROCESSING        BatchListEdgeAppsRequestFunctionType
	PROTOCOL_PARSING       BatchListEdgeAppsRequestFunctionType
	ON_PREMISE_INTEGRATION BatchListEdgeAppsRequestFunctionType
	GATEWAY_MANAGER        BatchListEdgeAppsRequestFunctionType
	COMPOSITE_APPLICATION  BatchListEdgeAppsRequestFunctionType
	DATA_COLLECTION        BatchListEdgeAppsRequestFunctionType
}

func GetBatchListEdgeAppsRequestFunctionTypeEnum() BatchListEdgeAppsRequestFunctionTypeEnum {
	return BatchListEdgeAppsRequestFunctionTypeEnum{
		DATA_PROCESSING: BatchListEdgeAppsRequestFunctionType{
			value: "DATA_PROCESSING",
		},
		PROTOCOL_PARSING: BatchListEdgeAppsRequestFunctionType{
			value: "PROTOCOL_PARSING",
		},
		ON_PREMISE_INTEGRATION: BatchListEdgeAppsRequestFunctionType{
			value: "ON_PREMISE_INTEGRATION",
		},
		GATEWAY_MANAGER: BatchListEdgeAppsRequestFunctionType{
			value: "GATEWAY_MANAGER",
		},
		COMPOSITE_APPLICATION: BatchListEdgeAppsRequestFunctionType{
			value: "COMPOSITE_APPLICATION",
		},
		DATA_COLLECTION: BatchListEdgeAppsRequestFunctionType{
			value: "DATA_COLLECTION",
		},
	}
}

func (c BatchListEdgeAppsRequestFunctionType) Value() string {
	return c.value
}

func (c BatchListEdgeAppsRequestFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListEdgeAppsRequestFunctionType) UnmarshalJSON(b []byte) error {
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
