package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DescribeBackupEncryptStatusResponse Response Object
type DescribeBackupEncryptStatusResponse struct {

	// 备份加密状态。
	EncryptionStatus *DescribeBackupEncryptStatusResponseEncryptionStatus `json:"encryption_status,omitempty"`
	HttpStatusCode   int                                                  `json:"-"`
}

func (o DescribeBackupEncryptStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeBackupEncryptStatusResponse struct{}"
	}

	return strings.Join([]string{"DescribeBackupEncryptStatusResponse", string(data)}, " ")
}

type DescribeBackupEncryptStatusResponseEncryptionStatus struct {
	value string
}

type DescribeBackupEncryptStatusResponseEncryptionStatusEnum struct {
	ON  DescribeBackupEncryptStatusResponseEncryptionStatus
	OFF DescribeBackupEncryptStatusResponseEncryptionStatus
}

func GetDescribeBackupEncryptStatusResponseEncryptionStatusEnum() DescribeBackupEncryptStatusResponseEncryptionStatusEnum {
	return DescribeBackupEncryptStatusResponseEncryptionStatusEnum{
		ON: DescribeBackupEncryptStatusResponseEncryptionStatus{
			value: "on",
		},
		OFF: DescribeBackupEncryptStatusResponseEncryptionStatus{
			value: "off",
		},
	}
}

func (c DescribeBackupEncryptStatusResponseEncryptionStatus) Value() string {
	return c.value
}

func (c DescribeBackupEncryptStatusResponseEncryptionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DescribeBackupEncryptStatusResponseEncryptionStatus) UnmarshalJSON(b []byte) error {
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
