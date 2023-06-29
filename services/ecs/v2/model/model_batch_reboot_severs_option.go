package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchRebootSeversOption
type BatchRebootSeversOption struct {

	// 云服务器ID列表。
	Servers []ServerId `json:"servers"`

	// 重启类型：  - SOFT：普通重启。 - HARD：强制重启。
	Type BatchRebootSeversOptionType `json:"type"`
}

func (o BatchRebootSeversOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootSeversOption struct{}"
	}

	return strings.Join([]string{"BatchRebootSeversOption", string(data)}, " ")
}

type BatchRebootSeversOptionType struct {
	value string
}

type BatchRebootSeversOptionTypeEnum struct {
	SOFT BatchRebootSeversOptionType
	HARD BatchRebootSeversOptionType
}

func GetBatchRebootSeversOptionTypeEnum() BatchRebootSeversOptionTypeEnum {
	return BatchRebootSeversOptionTypeEnum{
		SOFT: BatchRebootSeversOptionType{
			value: "SOFT",
		},
		HARD: BatchRebootSeversOptionType{
			value: "HARD",
		},
	}
}

func (c BatchRebootSeversOptionType) Value() string {
	return c.value
}

func (c BatchRebootSeversOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRebootSeversOptionType) UnmarshalJSON(b []byte) error {
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
