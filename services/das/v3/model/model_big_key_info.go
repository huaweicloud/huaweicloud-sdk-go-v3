package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BigKeyInfo struct {

	// key名称
	Name *string `json:"name,omitempty"`

	// key类型。取值：string、hash、list、zset、set、stream、exhash
	Type *BigKeyInfoType `json:"type,omitempty"`

	// 长度或存储占用值
	Size *int64 `json:"size,omitempty"`

	// DB ID
	DbId *string `json:"db_id,omitempty"`
}

func (o BigKeyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BigKeyInfo struct{}"
	}

	return strings.Join([]string{"BigKeyInfo", string(data)}, " ")
}

type BigKeyInfoType struct {
	value string
}

type BigKeyInfoTypeEnum struct {
	STRING BigKeyInfoType
	HASH   BigKeyInfoType
	LIST   BigKeyInfoType
	ZSET   BigKeyInfoType
	SET    BigKeyInfoType
	STREAM BigKeyInfoType
	EXHASH BigKeyInfoType
}

func GetBigKeyInfoTypeEnum() BigKeyInfoTypeEnum {
	return BigKeyInfoTypeEnum{
		STRING: BigKeyInfoType{
			value: "string",
		},
		HASH: BigKeyInfoType{
			value: "hash",
		},
		LIST: BigKeyInfoType{
			value: "list",
		},
		ZSET: BigKeyInfoType{
			value: "zset",
		},
		SET: BigKeyInfoType{
			value: "set",
		},
		STREAM: BigKeyInfoType{
			value: "stream",
		},
		EXHASH: BigKeyInfoType{
			value: "exhash",
		},
	}
}

func (c BigKeyInfoType) Value() string {
	return c.value
}

func (c BigKeyInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BigKeyInfoType) UnmarshalJSON(b []byte) error {
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
