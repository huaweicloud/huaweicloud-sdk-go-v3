package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SupabaseHaConfig HA配置。
type SupabaseHaConfig struct {
	Mode *SupabaseHaConfigMode `json:"mode,omitempty"`

	ReplicationMode *SupabaseHaConfigReplicationMode `json:"replication_mode,omitempty"`
}

func (o SupabaseHaConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupabaseHaConfig struct{}"
	}

	return strings.Join([]string{"SupabaseHaConfig", string(data)}, " ")
}

type SupabaseHaConfigMode struct {
	value string
}

type SupabaseHaConfigModeEnum struct {
	HA     SupabaseHaConfigMode
	SINGLE SupabaseHaConfigMode
}

func GetSupabaseHaConfigModeEnum() SupabaseHaConfigModeEnum {
	return SupabaseHaConfigModeEnum{
		HA: SupabaseHaConfigMode{
			value: "Ha",
		},
		SINGLE: SupabaseHaConfigMode{
			value: "Single",
		},
	}
}

func (c SupabaseHaConfigMode) Value() string {
	return c.value
}

func (c SupabaseHaConfigMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SupabaseHaConfigMode) UnmarshalJSON(b []byte) error {
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

type SupabaseHaConfigReplicationMode struct {
	value string
}

type SupabaseHaConfigReplicationModeEnum struct {
	ASYNC SupabaseHaConfigReplicationMode
	SYNC  SupabaseHaConfigReplicationMode
}

func GetSupabaseHaConfigReplicationModeEnum() SupabaseHaConfigReplicationModeEnum {
	return SupabaseHaConfigReplicationModeEnum{
		ASYNC: SupabaseHaConfigReplicationMode{
			value: "async",
		},
		SYNC: SupabaseHaConfigReplicationMode{
			value: "sync",
		},
	}
}

func (c SupabaseHaConfigReplicationMode) Value() string {
	return c.value
}

func (c SupabaseHaConfigReplicationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SupabaseHaConfigReplicationMode) UnmarshalJSON(b []byte) error {
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
