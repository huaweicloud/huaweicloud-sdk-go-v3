package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDatasourceColumnsRequest Request Object
type ListDatasourceColumnsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 数据源ID
	DatasourceId string `json:"datasource_id"`

	// 数据源所在任务位置 - SOURCE 数据源处于任务源端 - TARGET 数据源处于任务目标端
	Position ListDatasourceColumnsRequestPosition `json:"position"`

	// 数据库名称，只支持MRSHIVE，FIHIVE类型的数据源
	DbName *string `json:"db_name,omitempty"`

	// 字段所在的表名
	TableName *string `json:"table_name,omitempty"`
}

func (o ListDatasourceColumnsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasourceColumnsRequest struct{}"
	}

	return strings.Join([]string{"ListDatasourceColumnsRequest", string(data)}, " ")
}

type ListDatasourceColumnsRequestPosition struct {
	value string
}

type ListDatasourceColumnsRequestPositionEnum struct {
	SOURCE ListDatasourceColumnsRequestPosition
	TARGET ListDatasourceColumnsRequestPosition
}

func GetListDatasourceColumnsRequestPositionEnum() ListDatasourceColumnsRequestPositionEnum {
	return ListDatasourceColumnsRequestPositionEnum{
		SOURCE: ListDatasourceColumnsRequestPosition{
			value: "SOURCE",
		},
		TARGET: ListDatasourceColumnsRequestPosition{
			value: "TARGET",
		},
	}
}

func (c ListDatasourceColumnsRequestPosition) Value() string {
	return c.value
}

func (c ListDatasourceColumnsRequestPosition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDatasourceColumnsRequestPosition) UnmarshalJSON(b []byte) error {
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
