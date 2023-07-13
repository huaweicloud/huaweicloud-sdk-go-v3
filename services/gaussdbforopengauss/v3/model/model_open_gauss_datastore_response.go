package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussDatastoreResponse 数据库版本。不填时，默认为当前最新版本。  数据库支持版本可根据查询数据库引擎版本接口获取
type OpenGaussDatastoreResponse struct {

	// 数据库引擎，不区分大小写，取值如下：  GaussDB(for openGauss)
	Type OpenGaussDatastoreResponseType `json:"type"`

	// 数据库版本。
	Version string `json:"version"`
}

func (o OpenGaussDatastoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussDatastoreResponse struct{}"
	}

	return strings.Join([]string{"OpenGaussDatastoreResponse", string(data)}, " ")
}

type OpenGaussDatastoreResponseType struct {
	value string
}

type OpenGaussDatastoreResponseTypeEnum struct {
	GAUSS_DB                OpenGaussDatastoreResponseType
	GAUSS_DB_FOR_OPEN_GAUSS OpenGaussDatastoreResponseType
}

func GetOpenGaussDatastoreResponseTypeEnum() OpenGaussDatastoreResponseTypeEnum {
	return OpenGaussDatastoreResponseTypeEnum{
		GAUSS_DB: OpenGaussDatastoreResponseType{
			value: "GaussDB",
		},
		GAUSS_DB_FOR_OPEN_GAUSS: OpenGaussDatastoreResponseType{
			value: "GaussDB(for openGauss)",
		},
	}
}

func (c OpenGaussDatastoreResponseType) Value() string {
	return c.value
}

func (c OpenGaussDatastoreResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussDatastoreResponseType) UnmarshalJSON(b []byte) error {
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
