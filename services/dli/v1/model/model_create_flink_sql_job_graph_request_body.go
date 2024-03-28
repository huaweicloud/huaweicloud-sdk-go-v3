package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateFlinkSqlJobGraphRequestBody 生成SQL静态流图请求参数
type CreateFlinkSqlJobGraphRequestBody struct {

	// SQL
	SqlBody string `json:"sql_body"`

	// Flink版本。当前只支持1.10和1.12。
	FlinkVersion *CreateFlinkSqlJobGraphRequestBodyFlinkVersion `json:"flink_version,omitempty"`

	// CU总数
	CuNumber *int32 `json:"cu_number,omitempty"`

	// 管理单元CU数量
	ManagerCuNumber *int32 `json:"manager_cu_number,omitempty"`

	// 最大并行度
	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	// 单个taskManagerCU数量
	TmCus *int32 `json:"tm_cus,omitempty"`

	// 单个taskManager Slot数量
	TmSlotNum *int32 `json:"tm_slot_num,omitempty"`

	// 算子的配置
	OperatorConfig *string `json:"operator_config,omitempty"`

	// 是否静态资源预估
	StaticEstimator *bool `json:"static_estimator,omitempty"`

	// 作业类型
	JobType *string `json:"job_type,omitempty"`

	// 流图类型。当前支持以下两种流图类型。 简化流图：simple_graph 静态流图：job_graph。
	GraphType *CreateFlinkSqlJobGraphRequestBodyGraphType `json:"graph_type,omitempty"`

	// 每个算子的流量/命中率配置，json格式的字符串。
	StaticEstimatorConfig *string `json:"static_estimator_config,omitempty"`
}

func (o CreateFlinkSqlJobGraphRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSqlJobGraphRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlinkSqlJobGraphRequestBody", string(data)}, " ")
}

type CreateFlinkSqlJobGraphRequestBodyFlinkVersion struct {
	value string
}

type CreateFlinkSqlJobGraphRequestBodyFlinkVersionEnum struct {
	E_1_1  CreateFlinkSqlJobGraphRequestBodyFlinkVersion
	E_1_12 CreateFlinkSqlJobGraphRequestBodyFlinkVersion
}

func GetCreateFlinkSqlJobGraphRequestBodyFlinkVersionEnum() CreateFlinkSqlJobGraphRequestBodyFlinkVersionEnum {
	return CreateFlinkSqlJobGraphRequestBodyFlinkVersionEnum{
		E_1_1: CreateFlinkSqlJobGraphRequestBodyFlinkVersion{
			value: "1.1",
		},
		E_1_12: CreateFlinkSqlJobGraphRequestBodyFlinkVersion{
			value: "1.12",
		},
	}
}

func (c CreateFlinkSqlJobGraphRequestBodyFlinkVersion) Value() string {
	return c.value
}

func (c CreateFlinkSqlJobGraphRequestBodyFlinkVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFlinkSqlJobGraphRequestBodyFlinkVersion) UnmarshalJSON(b []byte) error {
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

type CreateFlinkSqlJobGraphRequestBodyGraphType struct {
	value string
}

type CreateFlinkSqlJobGraphRequestBodyGraphTypeEnum struct {
	SIMPLE_GRAPH CreateFlinkSqlJobGraphRequestBodyGraphType
	JOB_GRAPH    CreateFlinkSqlJobGraphRequestBodyGraphType
}

func GetCreateFlinkSqlJobGraphRequestBodyGraphTypeEnum() CreateFlinkSqlJobGraphRequestBodyGraphTypeEnum {
	return CreateFlinkSqlJobGraphRequestBodyGraphTypeEnum{
		SIMPLE_GRAPH: CreateFlinkSqlJobGraphRequestBodyGraphType{
			value: "simple_graph",
		},
		JOB_GRAPH: CreateFlinkSqlJobGraphRequestBodyGraphType{
			value: "job_graph",
		},
	}
}

func (c CreateFlinkSqlJobGraphRequestBodyGraphType) Value() string {
	return c.value
}

func (c CreateFlinkSqlJobGraphRequestBodyGraphType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFlinkSqlJobGraphRequestBodyGraphType) UnmarshalJSON(b []byte) error {
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
