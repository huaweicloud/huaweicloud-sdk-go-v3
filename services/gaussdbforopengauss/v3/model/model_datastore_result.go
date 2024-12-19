package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DatastoreResult 数据库版本。不填时，默认为当前最新版本。  数据库支持版本可根据查询数据库引擎版本接口获取
type DatastoreResult struct {

	// 数据库引擎，不区分大小写，取值如下：  GaussDB
	Type DatastoreResultType `json:"type"`

	// 数据库版本。
	Version string `json:"version"`
}

func (o DatastoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatastoreResult struct{}"
	}

	return strings.Join([]string{"DatastoreResult", string(data)}, " ")
}

type DatastoreResultType struct {
	value string
}

type DatastoreResultTypeEnum struct {
	GAUSS_DB DatastoreResultType
}

func GetDatastoreResultTypeEnum() DatastoreResultTypeEnum {
	return DatastoreResultTypeEnum{
		GAUSS_DB: DatastoreResultType{
			value: "GaussDB",
		},
	}
}

func (c DatastoreResultType) Value() string {
	return c.value
}

func (c DatastoreResultType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DatastoreResultType) UnmarshalJSON(b []byte) error {
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
