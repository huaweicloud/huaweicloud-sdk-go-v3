package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussDatastoreOption 数据库信息。
type OpenGaussDatastoreOption struct {

	// 数据库引擎，不区分大小写，取值如下：  GaussDB。
	Type OpenGaussDatastoreOptionType `json:"type"`

	// 数据库版本。不填时，默认为当前最新版本。  GaussDB支持的版本参考[查询数据库引擎的版本](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=GaussDB&api=ListDatastores)。
	Version *string `json:"version,omitempty"`
}

func (o OpenGaussDatastoreOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussDatastoreOption struct{}"
	}

	return strings.Join([]string{"OpenGaussDatastoreOption", string(data)}, " ")
}

type OpenGaussDatastoreOptionType struct {
	value string
}

type OpenGaussDatastoreOptionTypeEnum struct {
	GAUSS_DB OpenGaussDatastoreOptionType
}

func GetOpenGaussDatastoreOptionTypeEnum() OpenGaussDatastoreOptionTypeEnum {
	return OpenGaussDatastoreOptionTypeEnum{
		GAUSS_DB: OpenGaussDatastoreOptionType{
			value: "GaussDB",
		},
	}
}

func (c OpenGaussDatastoreOptionType) Value() string {
	return c.value
}

func (c OpenGaussDatastoreOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussDatastoreOptionType) UnmarshalJSON(b []byte) error {
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
