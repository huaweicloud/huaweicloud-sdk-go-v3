package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IcebergSortField 排序顺序字段，用于定义排序规则。
type IcebergSortField struct {

	// 源字段的id。
	SourceId int32 `json:"source_id"`

	// 转换函数。
	Transform string `json:"transform"`

	// 排序方向。
	Direction IcebergSortFieldDirection `json:"direction"`

	// null值的排序。
	NullOrder IcebergSortFieldNullOrder `json:"null_order"`
}

func (o IcebergSortField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IcebergSortField struct{}"
	}

	return strings.Join([]string{"IcebergSortField", string(data)}, " ")
}

type IcebergSortFieldDirection struct {
	value string
}

type IcebergSortFieldDirectionEnum struct {
	ASC  IcebergSortFieldDirection
	DESC IcebergSortFieldDirection
}

func GetIcebergSortFieldDirectionEnum() IcebergSortFieldDirectionEnum {
	return IcebergSortFieldDirectionEnum{
		ASC: IcebergSortFieldDirection{
			value: "ASC",
		},
		DESC: IcebergSortFieldDirection{
			value: "DESC",
		},
	}
}

func (c IcebergSortFieldDirection) Value() string {
	return c.value
}

func (c IcebergSortFieldDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IcebergSortFieldDirection) UnmarshalJSON(b []byte) error {
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

type IcebergSortFieldNullOrder struct {
	value string
}

type IcebergSortFieldNullOrderEnum struct {
	NULLS_FIRST IcebergSortFieldNullOrder
	NULLS_LAST  IcebergSortFieldNullOrder
}

func GetIcebergSortFieldNullOrderEnum() IcebergSortFieldNullOrderEnum {
	return IcebergSortFieldNullOrderEnum{
		NULLS_FIRST: IcebergSortFieldNullOrder{
			value: "NULLS_FIRST",
		},
		NULLS_LAST: IcebergSortFieldNullOrder{
			value: "NULLS_LAST",
		},
	}
}

func (c IcebergSortFieldNullOrder) Value() string {
	return c.value
}

func (c IcebergSortFieldNullOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IcebergSortFieldNullOrder) UnmarshalJSON(b []byte) error {
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
