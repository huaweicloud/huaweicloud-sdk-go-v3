package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TransferPolicyParam 转储策略信息
type TransferPolicyParam struct {

	// 转储目标前缀
	Prefix *string `json:"prefix,omitempty"`

	// 转储目标桶
	Destination string `json:"destination"`

	// 自动转储周期
	Period string `json:"period"`

	// 转储备份类型， db:自动备份 log:增量备份, snapshot:手动备份 snapshot:手动备份 db:自动备份 log:日志备份
	BackupType TransferPolicyParamBackupType `json:"backup_type"`
}

func (o TransferPolicyParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferPolicyParam struct{}"
	}

	return strings.Join([]string{"TransferPolicyParam", string(data)}, " ")
}

type TransferPolicyParamBackupType struct {
	value string
}

type TransferPolicyParamBackupTypeEnum struct {
	DB  TransferPolicyParamBackupType
	LOG TransferPolicyParamBackupType
}

func GetTransferPolicyParamBackupTypeEnum() TransferPolicyParamBackupTypeEnum {
	return TransferPolicyParamBackupTypeEnum{
		DB: TransferPolicyParamBackupType{
			value: "db",
		},
		LOG: TransferPolicyParamBackupType{
			value: "log",
		},
	}
}

func (c TransferPolicyParamBackupType) Value() string {
	return c.value
}

func (c TransferPolicyParamBackupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransferPolicyParamBackupType) UnmarshalJSON(b []byte) error {
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
