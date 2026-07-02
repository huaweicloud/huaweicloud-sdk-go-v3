package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyBackupVaultLockResponse Response Object
type ModifyBackupVaultLockResponse struct {

	// **参数解释**：  备份合规锁开启或延期状态。  **取值范围**： - success：成功。
	Result         *ModifyBackupVaultLockResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ModifyBackupVaultLockResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyBackupVaultLockResponse struct{}"
	}

	return strings.Join([]string{"ModifyBackupVaultLockResponse", string(data)}, " ")
}

type ModifyBackupVaultLockResponseResult struct {
	value string
}

type ModifyBackupVaultLockResponseResultEnum struct {
	SUCCESS ModifyBackupVaultLockResponseResult
}

func GetModifyBackupVaultLockResponseResultEnum() ModifyBackupVaultLockResponseResultEnum {
	return ModifyBackupVaultLockResponseResultEnum{
		SUCCESS: ModifyBackupVaultLockResponseResult{
			value: "success",
		},
	}
}

func (c ModifyBackupVaultLockResponseResult) Value() string {
	return c.value
}

func (c ModifyBackupVaultLockResponseResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyBackupVaultLockResponseResult) UnmarshalJSON(b []byte) error {
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
