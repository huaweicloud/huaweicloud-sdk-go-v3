package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifyBackupEncryptStatusResponse Response Object
type ModifyBackupEncryptStatusResponse struct {

	// 备份加密状态。
	EncryptionStatus *ModifyBackupEncryptStatusResponseEncryptionStatus `json:"encryption_status,omitempty"`
	HttpStatusCode   int                                                `json:"-"`
}

func (o ModifyBackupEncryptStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyBackupEncryptStatusResponse struct{}"
	}

	return strings.Join([]string{"ModifyBackupEncryptStatusResponse", string(data)}, " ")
}

type ModifyBackupEncryptStatusResponseEncryptionStatus struct {
	value string
}

type ModifyBackupEncryptStatusResponseEncryptionStatusEnum struct {
	ON  ModifyBackupEncryptStatusResponseEncryptionStatus
	OFF ModifyBackupEncryptStatusResponseEncryptionStatus
}

func GetModifyBackupEncryptStatusResponseEncryptionStatusEnum() ModifyBackupEncryptStatusResponseEncryptionStatusEnum {
	return ModifyBackupEncryptStatusResponseEncryptionStatusEnum{
		ON: ModifyBackupEncryptStatusResponseEncryptionStatus{
			value: "on",
		},
		OFF: ModifyBackupEncryptStatusResponseEncryptionStatus{
			value: "off",
		},
	}
}

func (c ModifyBackupEncryptStatusResponseEncryptionStatus) Value() string {
	return c.value
}

func (c ModifyBackupEncryptStatusResponseEncryptionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyBackupEncryptStatusResponseEncryptionStatus) UnmarshalJSON(b []byte) error {
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
