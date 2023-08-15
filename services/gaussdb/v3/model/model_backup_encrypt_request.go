package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackupEncryptRequest taurus支持备份加密的请求体。
type BackupEncryptRequest struct {

	// 加密类型。当前只支持kms。 开启加密时必传，关闭加密时不传。 不区分大小写。
	Type *BackupEncryptRequestType `json:"type,omitempty"`

	// kms加密ID。加密时必传，关闭加密时候不传。
	KmsKeyId *string `json:"kms_key_id,omitempty"`

	// 开启或关闭加密。不区分大小写。
	EncryptionStatus BackupEncryptRequestEncryptionStatus `json:"encryption_status"`
}

func (o BackupEncryptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupEncryptRequest struct{}"
	}

	return strings.Join([]string{"BackupEncryptRequest", string(data)}, " ")
}

type BackupEncryptRequestType struct {
	value string
}

type BackupEncryptRequestTypeEnum struct {
	KMS BackupEncryptRequestType
}

func GetBackupEncryptRequestTypeEnum() BackupEncryptRequestTypeEnum {
	return BackupEncryptRequestTypeEnum{
		KMS: BackupEncryptRequestType{
			value: "kms",
		},
	}
}

func (c BackupEncryptRequestType) Value() string {
	return c.value
}

func (c BackupEncryptRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupEncryptRequestType) UnmarshalJSON(b []byte) error {
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

type BackupEncryptRequestEncryptionStatus struct {
	value string
}

type BackupEncryptRequestEncryptionStatusEnum struct {
	ON  BackupEncryptRequestEncryptionStatus
	OFF BackupEncryptRequestEncryptionStatus
}

func GetBackupEncryptRequestEncryptionStatusEnum() BackupEncryptRequestEncryptionStatusEnum {
	return BackupEncryptRequestEncryptionStatusEnum{
		ON: BackupEncryptRequestEncryptionStatus{
			value: "on",
		},
		OFF: BackupEncryptRequestEncryptionStatus{
			value: "off",
		},
	}
}

func (c BackupEncryptRequestEncryptionStatus) Value() string {
	return c.value
}

func (c BackupEncryptRequestEncryptionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupEncryptRequestEncryptionStatus) UnmarshalJSON(b []byte) error {
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
