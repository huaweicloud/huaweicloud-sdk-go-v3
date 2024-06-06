package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchSgcComputeDimensionsRequest Request Object
type SearchSgcComputeDimensionsRequest struct {

	// DataArts Studio实例ID。
	InstanceId string `json:"instance_id"`

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	WorkspaceId string `json:"workspace_id"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 返回数据条数，表示页大小
	Limit *int32 `json:"limit,omitempty"`

	// 节点实例名称或脚本名称
	Name *string `json:"name,omitempty"`

	// 节点实例创建者name
	NodeCreatorName *string `json:"node_creator_name,omitempty"`

	// 类型，0表示节点实例,1表示脚本,2表示节点的测试运行,成本管理页面对于0和2的情况均展示为节点实例
	Type *SearchSgcComputeDimensionsRequestType `json:"type,omitempty"`

	// 节点类型，HIVE SQL, SparkSQL, Spark, Flink SQL, MRS Flink Job, DWS SQL为支持的枚举节点类型
	NodeType *SearchSgcComputeDimensionsRequestNodeType `json:"node_type,omitempty"`

	// 排序规则，示例priority ASC表示按照优先级升序排序,priority DESC表示按照优先级降序排序
	OrderBy *string `json:"order_by,omitempty"`
}

func (o SearchSgcComputeDimensionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSgcComputeDimensionsRequest struct{}"
	}

	return strings.Join([]string{"SearchSgcComputeDimensionsRequest", string(data)}, " ")
}

type SearchSgcComputeDimensionsRequestType struct {
	value int32
}

type SearchSgcComputeDimensionsRequestTypeEnum struct {
	E_0 SearchSgcComputeDimensionsRequestType
	E_1 SearchSgcComputeDimensionsRequestType
	E_2 SearchSgcComputeDimensionsRequestType
}

func GetSearchSgcComputeDimensionsRequestTypeEnum() SearchSgcComputeDimensionsRequestTypeEnum {
	return SearchSgcComputeDimensionsRequestTypeEnum{
		E_0: SearchSgcComputeDimensionsRequestType{
			value: 0,
		}, E_1: SearchSgcComputeDimensionsRequestType{
			value: 1,
		}, E_2: SearchSgcComputeDimensionsRequestType{
			value: 2,
		},
	}
}

func (c SearchSgcComputeDimensionsRequestType) Value() int32 {
	return c.value
}

func (c SearchSgcComputeDimensionsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchSgcComputeDimensionsRequestType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type SearchSgcComputeDimensionsRequestNodeType struct {
	value string
}

type SearchSgcComputeDimensionsRequestNodeTypeEnum struct {
	HIVE_SQL      SearchSgcComputeDimensionsRequestNodeType
	SPARK_SQL     SearchSgcComputeDimensionsRequestNodeType
	SPARK         SearchSgcComputeDimensionsRequestNodeType
	FLINK_SQL     SearchSgcComputeDimensionsRequestNodeType
	MRS_FLINK_JOB SearchSgcComputeDimensionsRequestNodeType
	DWS_SQL       SearchSgcComputeDimensionsRequestNodeType
}

func GetSearchSgcComputeDimensionsRequestNodeTypeEnum() SearchSgcComputeDimensionsRequestNodeTypeEnum {
	return SearchSgcComputeDimensionsRequestNodeTypeEnum{
		HIVE_SQL: SearchSgcComputeDimensionsRequestNodeType{
			value: "HIVE SQL",
		},
		SPARK_SQL: SearchSgcComputeDimensionsRequestNodeType{
			value: "SparkSQL",
		},
		SPARK: SearchSgcComputeDimensionsRequestNodeType{
			value: "Spark",
		},
		FLINK_SQL: SearchSgcComputeDimensionsRequestNodeType{
			value: "Flink SQL",
		},
		MRS_FLINK_JOB: SearchSgcComputeDimensionsRequestNodeType{
			value: "MRS Flink Job",
		},
		DWS_SQL: SearchSgcComputeDimensionsRequestNodeType{
			value: "DWS SQL",
		},
	}
}

func (c SearchSgcComputeDimensionsRequestNodeType) Value() string {
	return c.value
}

func (c SearchSgcComputeDimensionsRequestNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchSgcComputeDimensionsRequestNodeType) UnmarshalJSON(b []byte) error {
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
