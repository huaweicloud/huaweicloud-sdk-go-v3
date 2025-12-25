package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDataclassTypeRequestBody 创建类型请求参数
type CreateDataclassTypeRequestBody struct {

	// 子类型名称，告警类型、事件类型必填；威胁情报、漏洞类型、自定义类型不填
	SubCategory *string `json:"sub_category,omitempty"`

	// 子类型标识。告警类型、事件类型必填；威胁情报、漏洞类型、自定义类型不填
	SubCategoryCode *string `json:"sub_category_code,omitempty"`

	// 类型名称，告警类型、事件类型、威胁情报、漏洞类型必填；自定义类型不填
	Category *string `json:"category,omitempty"`

	// 类型标识。告警类型、事件类型、威胁情报、漏洞类型必填；自定义类型不填
	CategoryCode *string `json:"category_code,omitempty"`

	// 自定义类型名称。创建类型时，代表类型名称，创建子类型时，代表子类型名称。（自定义类型必填，其余类型不填）
	Name *string `json:"name,omitempty"`

	// 自定义类型标识。创建类型时代表类型标识；创建子类型时代表子类型标识。（自定义类型必填，其余类型不填）
	BusinessCode *string `json:"business_code,omitempty"`

	// 数据类业务编码，创建自定义类型必填。（其余类型不填）
	DataclassBusinessCode *string `json:"dataclass_business_code,omitempty"`

	// 类型描述
	Description *string `json:"description,omitempty"`

	// 类型层级，创建自定义类型必填，枚举值：1-类型，2-子类型。（其余类型不填）
	Level *CreateDataclassTypeRequestBodyLevel `json:"level,omitempty"`

	// 类型启用禁用状态，启用：true，禁用：false
	Enabled bool `json:"enabled"`

	// 响应时长
	Sla int64 `json:"sla"`
}

func (o CreateDataclassTypeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataclassTypeRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDataclassTypeRequestBody", string(data)}, " ")
}

type CreateDataclassTypeRequestBodyLevel struct {
	value int32
}

type CreateDataclassTypeRequestBodyLevelEnum struct {
	E_1 CreateDataclassTypeRequestBodyLevel
	E_2 CreateDataclassTypeRequestBodyLevel
}

func GetCreateDataclassTypeRequestBodyLevelEnum() CreateDataclassTypeRequestBodyLevelEnum {
	return CreateDataclassTypeRequestBodyLevelEnum{
		E_1: CreateDataclassTypeRequestBodyLevel{
			value: 1,
		}, E_2: CreateDataclassTypeRequestBodyLevel{
			value: 2,
		},
	}
}

func (c CreateDataclassTypeRequestBodyLevel) Value() int32 {
	return c.value
}

func (c CreateDataclassTypeRequestBodyLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDataclassTypeRequestBodyLevel) UnmarshalJSON(b []byte) error {
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
