package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
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

func (c ListMqsInstanceRequestIncludeInternal) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMqsInstanceRequestIncludeInternal) UnmarshalJSON(b []byte) error {
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
