package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListMqsInstanceRequest Request Object
type ListMqsInstanceRequest struct {

	// 是否包含内部的实例。include_internal参数必须为true。
	IncludeInternal ListMqsInstanceRequestIncludeInternal `json:"include_internal"`
}

func (o ListMqsInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMqsInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListMqsInstanceRequest", string(data)}, " ")
}

type ListMqsInstanceRequestIncludeInternal struct {
	value string
}

type ListMqsInstanceRequestIncludeInternalEnum struct {
	TRUE ListMqsInstanceRequestIncludeInternal
}

func GetListMqsInstanceRequestIncludeInternalEnum() ListMqsInstanceRequestIncludeInternalEnum {
	return ListMqsInstanceRequestIncludeInternalEnum{
		TRUE: ListMqsInstanceRequestIncludeInternal{
			value: "true",
		},
	}
}

func (c ListMqsInstanceRequestIncludeInternal) Value() string {
	return c.value
}

func (c ListMqsInstanceRequestIncludeInternal) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMqsInstanceRequestIncludeInternal) UnmarshalJSON(b []byte) error {
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
