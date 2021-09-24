package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 数据库信息。
type OpenGaussDatastoreResponse struct {
	// 数据库引擎，返回值如下：  GaussDB(openGauss)

	Type OpenGaussDatastoreResponseType `json:"type"`
	// 数据库版本。

	Version string `json:"version"`
}

func (o OpenGaussDatastoreResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OpenGaussDatastoreResponse struct{}"
	}

	return strings.Join([]string{"OpenGaussDatastoreResponse", string(data)}, " ")
}

type OpenGaussDatastoreResponseType struct {
	value string
}

type OpenGaussDatastoreResponseTypeEnum struct {
	GAUSS_DB_OPEN_GAUSS OpenGaussDatastoreResponseType
}

func GetOpenGaussDatastoreResponseTypeEnum() OpenGaussDatastoreResponseTypeEnum {
	return OpenGaussDatastoreResponseTypeEnum{
		GAUSS_DB_OPEN_GAUSS: OpenGaussDatastoreResponseType{
			value: "GaussDB(openGauss)",
		},
	}
}

func (c OpenGaussDatastoreResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *OpenGaussDatastoreResponseType) UnmarshalJSON(b []byte) error {
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
