package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 日志转储信息
type UpdateTransferRequestBodyLogTransferInfo struct {
	// 日志转储格式。只支持\"RAW\", \"JSON\"。RAW是指原始日志格式，JSON是指JSON日志格式。OBS转储和DIS转储支持JSON和RAW，DMS转储仅支持RAW

	LogStorageFormat UpdateTransferRequestBodyLogTransferInfoLogStorageFormat `json:"log_storage_format"`
	// 日志转储状态，ENABLE是指日志转储开启状态，DISABLE是指日志转储关闭状态，EXCEPTION是指日志转储异常状态

	LogTransferStatus UpdateTransferRequestBodyLogTransferInfoLogTransferStatus `json:"log_transfer_status"`
	// 日志转储详细信息

	LogTransferDetail *TransferDetail `json:"log_transfer_detail"`
}

func (o UpdateTransferRequestBodyLogTransferInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTransferRequestBodyLogTransferInfo struct{}"
	}

	return strings.Join([]string{"UpdateTransferRequestBodyLogTransferInfo", string(data)}, " ")
}

type UpdateTransferRequestBodyLogTransferInfoLogStorageFormat struct {
	value string
}

type UpdateTransferRequestBodyLogTransferInfoLogStorageFormatEnum struct {
	JSON UpdateTransferRequestBodyLogTransferInfoLogStorageFormat
	RAW  UpdateTransferRequestBodyLogTransferInfoLogStorageFormat
}

func GetUpdateTransferRequestBodyLogTransferInfoLogStorageFormatEnum() UpdateTransferRequestBodyLogTransferInfoLogStorageFormatEnum {
	return UpdateTransferRequestBodyLogTransferInfoLogStorageFormatEnum{
		JSON: UpdateTransferRequestBodyLogTransferInfoLogStorageFormat{
			value: "JSON",
		},
		RAW: UpdateTransferRequestBodyLogTransferInfoLogStorageFormat{
			value: "RAW",
		},
	}
}

func (c UpdateTransferRequestBodyLogTransferInfoLogStorageFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTransferRequestBodyLogTransferInfoLogStorageFormat) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateTransferRequestBodyLogTransferInfoLogTransferStatus struct {
	value string
}

type UpdateTransferRequestBodyLogTransferInfoLogTransferStatusEnum struct {
	ENABLE    UpdateTransferRequestBodyLogTransferInfoLogTransferStatus
	DISABLE   UpdateTransferRequestBodyLogTransferInfoLogTransferStatus
	EXCEPTION UpdateTransferRequestBodyLogTransferInfoLogTransferStatus
}

func GetUpdateTransferRequestBodyLogTransferInfoLogTransferStatusEnum() UpdateTransferRequestBodyLogTransferInfoLogTransferStatusEnum {
	return UpdateTransferRequestBodyLogTransferInfoLogTransferStatusEnum{
		ENABLE: UpdateTransferRequestBodyLogTransferInfoLogTransferStatus{
			value: "ENABLE",
		},
		DISABLE: UpdateTransferRequestBodyLogTransferInfoLogTransferStatus{
			value: "DISABLE",
		},
		EXCEPTION: UpdateTransferRequestBodyLogTransferInfoLogTransferStatus{
			value: "EXCEPTION",
		},
	}
}

func (c UpdateTransferRequestBodyLogTransferInfoLogTransferStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTransferRequestBodyLogTransferInfoLogTransferStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
