package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussDatastoreResult 数据库版本。不填时，默认为当前最新版本。  数据库支持版本可根据查询数据库引擎版本接口获取
type OpenGaussDatastoreResult struct {

	// 数据库引擎，不区分大小写，取值如下：  GaussDB
	Type OpenGaussDatastoreResultType `json:"type"`

	// 数据库版本。
	Version string `json:"version"`
}

func (o OpenGaussDatastoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussDatastoreResult struct{}"
	}

	return strings.Join([]string{"OpenGaussDatastoreResult", string(data)}, " ")
}

type OpenGaussDatastoreResultType struct {
	value string
}

type OpenGaussDatastoreResultTypeEnum struct {
	GAUSS_DB OpenGaussDatastoreResultType
}

func GetOpenGaussDatastoreResultTypeEnum() OpenGaussDatastoreResultTypeEnum {
	return OpenGaussDatastoreResultTypeEnum{
		GAUSS_DB: OpenGaussDatastoreResultType{
			value: "GaussDB",
		},
	}
}

func (c OpenGaussDatastoreResultType) Value() string {
	return c.value
}

func (c OpenGaussDatastoreResultType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussDatastoreResultType) UnmarshalJSON(b []byte) error {
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
