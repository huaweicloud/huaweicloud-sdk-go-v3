package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackupVaultLockRequest **参数解释**：  打开或延期备份合规锁配置的请求体。  **约束限制**：  不涉及。
type BackupVaultLockRequest struct {
	BackupVaultlockInfo *BackupVaultLockInfoRequest `json:"backup_vaultlock_info"`

	// **参数解释**：  操作类型。  **约束限制**：  当前只支持开启和延期。不区分大小写。  **取值范围**： - enable：开启。 - extend：延期。  **默认取值**：  enable。
	Action BackupVaultLockRequestAction `json:"action"`
}

func (o BackupVaultLockRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupVaultLockRequest struct{}"
	}

	return strings.Join([]string{"BackupVaultLockRequest", string(data)}, " ")
}

type BackupVaultLockRequestAction struct {
	value string
}

type BackupVaultLockRequestActionEnum struct {
	ENABLE BackupVaultLockRequestAction
	EXTEND BackupVaultLockRequestAction
}

func GetBackupVaultLockRequestActionEnum() BackupVaultLockRequestActionEnum {
	return BackupVaultLockRequestActionEnum{
		ENABLE: BackupVaultLockRequestAction{
			value: "enable",
		},
		EXTEND: BackupVaultLockRequestAction{
			value: "extend",
		},
	}
}

func (c BackupVaultLockRequestAction) Value() string {
	return c.value
}

func (c BackupVaultLockRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupVaultLockRequestAction) UnmarshalJSON(b []byte) error {
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
