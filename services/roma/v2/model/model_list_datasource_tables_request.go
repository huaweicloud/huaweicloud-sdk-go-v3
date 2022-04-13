package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListDatasourceTablesRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 数据源ID

	DatasourceId string `json:"datasource_id"`
	// 数据源所在任务位置 - SOURCE 数据源处于任务源端 - TARGET 数据源处于任务目标端

	Position ListDatasourceTablesRequestPosition `json:"position"`
	// 数据库名称，只支持MRSHIVE，FIHIVE类型的数据源

	DbName *string `json:"db_name,omitempty"`
	// 数据库模式,GAUSS100数据库使用

	DbSchema *string `json:"db_schema,omitempty"`
	// 表名模糊匹配过滤器

	Filter *string `json:"filter,omitempty"`
}

func (o ListDatasourceTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasourceTablesRequest struct{}"
	}

	return strings.Join([]string{"ListDatasourceTablesRequest", string(data)}, " ")
}

type ListDatasourceTablesRequestPosition struct {
	value string
}

type ListDatasourceTablesRequestPositionEnum struct {
	SOURCE ListDatasourceTablesRequestPosition
	TARGET ListDatasourceTablesRequestPosition
}

func GetListDatasourceTablesRequestPositionEnum() ListDatasourceTablesRequestPositionEnum {
	return ListDatasourceTablesRequestPositionEnum{
		SOURCE: ListDatasourceTablesRequestPosition{
			value: "SOURCE",
		},
		TARGET: ListDatasourceTablesRequestPosition{
			value: "TARGET",
		},
	}
}

func (c ListDatasourceTablesRequestPosition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatasourceTablesRequestPosition) UnmarshalJSON(b []byte) error {
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
