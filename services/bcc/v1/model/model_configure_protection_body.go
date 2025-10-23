package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfigureProtectionBody 资源配置保护结构体
type ConfigureProtectionBody struct {

	// 资源ID列表
	Ids []string `json:"ids"`

	// 资源保护采用的方式。分为云备份cbr和数据库db
	ProtectEngine ConfigureProtectionBodyProtectEngine `json:"protect_engine"`

	LocalVaults *[]ConfigureVaultItem `json:"local_vaults,omitempty"`

	RemoteVaults *[]ConfigureVaultItem `json:"remote_vaults,omitempty"`

	DbBackupPolicy *DbBackupPolicyBody `json:"db_backup_policy,omitempty"`

	DbOffsiteBackupPolicy *DbOffsiteBackupPolicyBody `json:"db_offsite_backup_policy,omitempty"`
}

func (o ConfigureProtectionBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigureProtectionBody struct{}"
	}

	return strings.Join([]string{"ConfigureProtectionBody", string(data)}, " ")
}

type ConfigureProtectionBodyProtectEngine struct {
	value string
}

type ConfigureProtectionBodyProtectEngineEnum struct {
	CBR ConfigureProtectionBodyProtectEngine
	DB  ConfigureProtectionBodyProtectEngine
}

func GetConfigureProtectionBodyProtectEngineEnum() ConfigureProtectionBodyProtectEngineEnum {
	return ConfigureProtectionBodyProtectEngineEnum{
		CBR: ConfigureProtectionBodyProtectEngine{
			value: "cbr",
		},
		DB: ConfigureProtectionBodyProtectEngine{
			value: "db",
		},
	}
}

func (c ConfigureProtectionBodyProtectEngine) Value() string {
	return c.value
}

func (c ConfigureProtectionBodyProtectEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigureProtectionBodyProtectEngine) UnmarshalJSON(b []byte) error {
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
