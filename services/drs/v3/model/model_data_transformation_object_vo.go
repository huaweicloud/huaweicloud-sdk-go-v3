package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataTransformationObjectVo 数据加工对象
type DataTransformationObjectVo struct {

	// 数据库对象、数据库表名称和过滤类型名称，例如格式为db1-*-*-tb1-*-*---conditionFilter--。
	Id *string `json:"id,omitempty"`

	// contentConditionalFilter：普通行过滤配置； configConditionalFilter：高级行过滤配置。
	DataTransformationType *string `json:"data_transformation_type,omitempty"`

	// 库名称。
	DbName *string `json:"db_name,omitempty"`

	// schema名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名称。
	TableName *string `json:"table_name,omitempty"`
}

func (o DataTransformationObjectVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataTransformationObjectVo struct{}"
	}

	return strings.Join([]string{"DataTransformationObjectVo", string(data)}, " ")
}
