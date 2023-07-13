package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DatasourceOrderPara struct {

	// 排序参数名称
	Name *string `json:"name,omitempty"`

	// 对应的参数字段
	Field *string `json:"field,omitempty"`

	// 是否可选
	Optional *bool `json:"optional,omitempty"`

	// 排序方式
	Sort *DatasourceOrderParaSort `json:"sort,omitempty"`

	// 排序参数顺序
	Order *int32 `json:"order,omitempty"`
}

func (o DatasourceOrderPara) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasourceOrderPara struct{}"
	}

	return strings.Join([]string{"DatasourceOrderPara", string(data)}, " ")
}

type DatasourceOrderParaSort struct {
	value string
}

type DatasourceOrderParaSortEnum struct {
	ASC    DatasourceOrderParaSort
	DESC   DatasourceOrderParaSort
	CUSTOM DatasourceOrderParaSort
}

func GetDatasourceOrderParaSortEnum() DatasourceOrderParaSortEnum {
	return DatasourceOrderParaSortEnum{
		ASC: DatasourceOrderParaSort{
			value: "ASC",
		},
		DESC: DatasourceOrderParaSort{
			value: "DESC",
		},
		CUSTOM: DatasourceOrderParaSort{
			value: "CUSTOM",
		},
	}
}

func (c DatasourceOrderParaSort) Value() string {
	return c.value
}

func (c DatasourceOrderParaSort) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatasourceOrderParaSort) UnmarshalJSON(b []byte) error {
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
