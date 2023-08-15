package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Sbc struct {

	// 连接策略。 - DISABLED：已禁用 - AUTO_DISCONNECT：自动断开 - AUTO_LOCK：自动锁屏
	SbcAutomaticDisconnection *SbcSbcAutomaticDisconnection `json:"sbc_automatic_disconnection,omitempty"`

	SbcAutomaticDisconnectionOptions *SbcAutomaticDisconnectionOptions `json:"sbc_automatic_disconnection_options,omitempty"`
}

func (o Sbc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Sbc struct{}"
	}

	return strings.Join([]string{"Sbc", string(data)}, " ")
}

type SbcSbcAutomaticDisconnection struct {
	value string
}

type SbcSbcAutomaticDisconnectionEnum struct {
	DISABLED        SbcSbcAutomaticDisconnection
	AUTO_DISCONNECT SbcSbcAutomaticDisconnection
	AUTO_LOCK       SbcSbcAutomaticDisconnection
}

func GetSbcSbcAutomaticDisconnectionEnum() SbcSbcAutomaticDisconnectionEnum {
	return SbcSbcAutomaticDisconnectionEnum{
		DISABLED: SbcSbcAutomaticDisconnection{
			value: "DISABLED",
		},
		AUTO_DISCONNECT: SbcSbcAutomaticDisconnection{
			value: "AUTO_DISCONNECT",
		},
		AUTO_LOCK: SbcSbcAutomaticDisconnection{
			value: "AUTO_LOCK",
		},
	}
}

func (c SbcSbcAutomaticDisconnection) Value() string {
	return c.value
}

func (c SbcSbcAutomaticDisconnection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SbcSbcAutomaticDisconnection) UnmarshalJSON(b []byte) error {
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
