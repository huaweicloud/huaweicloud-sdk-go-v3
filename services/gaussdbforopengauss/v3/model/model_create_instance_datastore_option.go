package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceDatastoreOption 数据库信息。
type CreateInstanceDatastoreOption struct {

	// 数据库引擎，不区分大小写，取值如下：  GaussDB。
	Type CreateInstanceDatastoreOptionType `json:"type"`

	// 数据库版本。不填时，默认为当前最新版本。  GaussDB支持的版本参考[查询数据库引擎的版本](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=GaussDBforopenGauss&api=ListDatastores)。
	Version *string `json:"version,omitempty"`
}

func (o CreateInstanceDatastoreOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceDatastoreOption struct{}"
	}

	return strings.Join([]string{"CreateInstanceDatastoreOption", string(data)}, " ")
}

type CreateInstanceDatastoreOptionType struct {
	value string
}

type CreateInstanceDatastoreOptionTypeEnum struct {
	GAUSS_DB CreateInstanceDatastoreOptionType
}

func GetCreateInstanceDatastoreOptionTypeEnum() CreateInstanceDatastoreOptionTypeEnum {
	return CreateInstanceDatastoreOptionTypeEnum{
		GAUSS_DB: CreateInstanceDatastoreOptionType{
			value: "GaussDB",
		},
	}
}

func (c CreateInstanceDatastoreOptionType) Value() string {
	return c.value
}

func (c CreateInstanceDatastoreOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceDatastoreOptionType) UnmarshalJSON(b []byte) error {
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
