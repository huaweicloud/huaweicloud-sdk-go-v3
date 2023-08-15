package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTableMetaRequest Request Object
type ListTableMetaRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 数据库通配符
	DatabaseNamePattern *string `json:"database_name_pattern,omitempty"`

	// 表名通配符
	TableNamePattern *string `json:"table_name_pattern,omitempty"`

	// 查询的表类型
	TableTypes *[]ListTableMetaRequestTableTypes `json:"table_types,omitempty"`

	// 返回的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID
	Marker *string `json:"marker,omitempty"`

	// 是否查询上一页
	ReversePage *bool `json:"reverse_page,omitempty"`
}

func (o ListTableMetaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableMetaRequest struct{}"
	}

	return strings.Join([]string{"ListTableMetaRequest", string(data)}, " ")
}

type ListTableMetaRequestTableTypes struct {
	value string
}

type ListTableMetaRequestTableTypesEnum struct {
	MANAGED_TABLE     ListTableMetaRequestTableTypes
	EXTERNAL_TABLE    ListTableMetaRequestTableTypes
	VIRTUAL_VIEW      ListTableMetaRequestTableTypes
	MATERIALIZED_VIEW ListTableMetaRequestTableTypes
}

func GetListTableMetaRequestTableTypesEnum() ListTableMetaRequestTableTypesEnum {
	return ListTableMetaRequestTableTypesEnum{
		MANAGED_TABLE: ListTableMetaRequestTableTypes{
			value: "MANAGED_TABLE",
		},
		EXTERNAL_TABLE: ListTableMetaRequestTableTypes{
			value: "EXTERNAL_TABLE",
		},
		VIRTUAL_VIEW: ListTableMetaRequestTableTypes{
			value: "VIRTUAL_VIEW",
		},
		MATERIALIZED_VIEW: ListTableMetaRequestTableTypes{
			value: "MATERIALIZED_VIEW",
		},
	}
}

func (c ListTableMetaRequestTableTypes) Value() string {
	return c.value
}

func (c ListTableMetaRequestTableTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTableMetaRequestTableTypes) UnmarshalJSON(b []byte) error {
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
