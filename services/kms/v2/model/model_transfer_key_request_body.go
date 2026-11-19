package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TransferKeyRequestBody 转移密钥请求体
type TransferKeyRequestBody struct {

	// **参数解释：** 待转移的密钥ID **约束限制：** 满足正则表达式^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$ **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyId string `json:"key_id"`

	// **参数解释：** 目标密钥空间ID **约束限制：** 满足正则表达式^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$ **取值范围：** 不涉及 **默认取值：** 不涉及
	TargetKeyspaceId string `json:"target_keyspace_id"`

	// **参数解释：** 转移模式 **约束限制：** ASSIGN模式仅支持对称密钥，且密钥必须未属于任何密钥空间，目标密钥空间类型需与密钥所属keystore类型匹配； **取值范围：** - ASSIGN：将密钥分配给目标密钥空间 **默认取值：** ASSIGN
	TransferMode *TransferKeyRequestBodyTransferMode `json:"transfer_mode,omitempty"`
}

func (o TransferKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferKeyRequestBody struct{}"
	}

	return strings.Join([]string{"TransferKeyRequestBody", string(data)}, " ")
}

type TransferKeyRequestBodyTransferMode struct {
	value string
}

type TransferKeyRequestBodyTransferModeEnum struct {
	ASSIGN TransferKeyRequestBodyTransferMode
}

func GetTransferKeyRequestBodyTransferModeEnum() TransferKeyRequestBodyTransferModeEnum {
	return TransferKeyRequestBodyTransferModeEnum{
		ASSIGN: TransferKeyRequestBodyTransferMode{
			value: "ASSIGN",
		},
	}
}

func (c TransferKeyRequestBodyTransferMode) Value() string {
	return c.value
}

func (c TransferKeyRequestBodyTransferMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransferKeyRequestBodyTransferMode) UnmarshalJSON(b []byte) error {
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
