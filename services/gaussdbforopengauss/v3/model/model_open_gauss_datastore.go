package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussDatastore 数据库信息。
type OpenGaussDatastore struct {

	// 数据库引擎，不区分大小写，取值如下：  GaussDB(for openGauss)。
	Type OpenGaussDatastoreType `json:"type"`

	// 数据库版本。不填时，默认为当前最新版本。  GaussDB支持的版本参考[查询数据库引擎的版本](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=GaussDBforopenGauss&api=ListDatastores)。
	Version *string `json:"version,omitempty"`
}

func (o OpenGaussDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussDatastore struct{}"
	}

	return strings.Join([]string{"OpenGaussDatastore", string(data)}, " ")
}

type OpenGaussDatastoreType struct {
	value string
}

type OpenGaussDatastoreTypeEnum struct {
	GAUSS_DB_FOR_OPEN_GAUSS OpenGaussDatastoreType
	GAUSS_DB                OpenGaussDatastoreType
}

func GetOpenGaussDatastoreTypeEnum() OpenGaussDatastoreTypeEnum {
	return OpenGaussDatastoreTypeEnum{
		GAUSS_DB_FOR_OPEN_GAUSS: OpenGaussDatastoreType{
			value: "GaussDB(for openGauss)",
		},
		GAUSS_DB: OpenGaussDatastoreType{
			value: "GaussDB",
		},
	}
}

func (c OpenGaussDatastoreType) Value() string {
	return c.value
}

func (c OpenGaussDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussDatastoreType) UnmarshalJSON(b []byte) error {
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
