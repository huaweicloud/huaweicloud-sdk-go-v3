package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLtsLogPolicyRespondBodyInstanceDatastore 实例的数据库引擎和版本。
type ListLtsLogPolicyRespondBodyInstanceDatastore struct {

	// 数据库引擎，值为mongodb。
	Type *ListLtsLogPolicyRespondBodyInstanceDatastoreType `json:"type,omitempty"`

	// 数据库大版本。
	Version *string `json:"version,omitempty"`
}

func (o ListLtsLogPolicyRespondBodyInstanceDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsLogPolicyRespondBodyInstanceDatastore struct{}"
	}

	return strings.Join([]string{"ListLtsLogPolicyRespondBodyInstanceDatastore", string(data)}, " ")
}

type ListLtsLogPolicyRespondBodyInstanceDatastoreType struct {
	value string
}

type ListLtsLogPolicyRespondBodyInstanceDatastoreTypeEnum struct {
	MONGODB ListLtsLogPolicyRespondBodyInstanceDatastoreType
}

func GetListLtsLogPolicyRespondBodyInstanceDatastoreTypeEnum() ListLtsLogPolicyRespondBodyInstanceDatastoreTypeEnum {
	return ListLtsLogPolicyRespondBodyInstanceDatastoreTypeEnum{
		MONGODB: ListLtsLogPolicyRespondBodyInstanceDatastoreType{
			value: "mongodb",
		},
	}
}

func (c ListLtsLogPolicyRespondBodyInstanceDatastoreType) Value() string {
	return c.value
}

func (c ListLtsLogPolicyRespondBodyInstanceDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLtsLogPolicyRespondBodyInstanceDatastoreType) UnmarshalJSON(b []byte) error {
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
