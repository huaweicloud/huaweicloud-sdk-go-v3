package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 日志转储信息
type CreateTransferRequestBodyLogTransferInfo struct {
	// 日志转储类型。OBS指OBS日志转储，DIS指DIS日志转储，DMS指DMS日志转储

	LogTransferType string `json:"log_transfer_type"`
	// 日志转储方式。cycle是指周期性转储，realTime是指实时转储。OBS转储只支持\"cycle\"，DIS转储和DMS转储只支持\"realTime\"。

	LogTransferMode CreateTransferRequestBodyLogTransferInfoLogTransferMode `json:"log_transfer_mode"`
	// 日志转储格式。只支持\"RAW\", \"JSON\"。RAW是指原始日志格式，JSON是指JSON日志格式。OBS转储和DIS转储支持JSON和RAW，DMS转储仅支持RAW

	LogStorageFormat CreateTransferRequestBodyLogTransferInfoLogStorageFormat `json:"log_storage_format"`
	// 日志转储状态，只支持\"ENABLE\",\"DISABLE\",\"EXCEPTION\"。ENABLE是指日志转储开启状态，DISABLE是指日志转储关闭状态，EXCEPTION是指日志转储异常状态

	LogTransferStatus CreateTransferRequestBodyLogTransferInfoLogTransferStatus `json:"log_transfer_status"`

	LogAgencyTransfer *CreateTransferRequestBodyLogTransferInfoLogAgencyTransfer `json:"log_agency_transfer,omitempty"`
	// 日志转储详细信息

	LogTransferDetail *TransferDetail `json:"log_transfer_detail"`
}

func (o CreateTransferRequestBodyLogTransferInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferRequestBodyLogTransferInfo struct{}"
	}

	return strings.Join([]string{"CreateTransferRequestBodyLogTransferInfo", string(data)}, " ")
}

type CreateTransferRequestBodyLogTransferInfoLogTransferMode struct {
	value string
}

type CreateTransferRequestBodyLogTransferInfoLogTransferModeEnum struct {
	CYCLE     CreateTransferRequestBodyLogTransferInfoLogTransferMode
	REAL_TIME CreateTransferRequestBodyLogTransferInfoLogTransferMode
}

func GetCreateTransferRequestBodyLogTransferInfoLogTransferModeEnum() CreateTransferRequestBodyLogTransferInfoLogTransferModeEnum {
	return CreateTransferRequestBodyLogTransferInfoLogTransferModeEnum{
		CYCLE: CreateTransferRequestBodyLogTransferInfoLogTransferMode{
			value: "cycle",
		},
		REAL_TIME: CreateTransferRequestBodyLogTransferInfoLogTransferMode{
			value: "realTime",
		},
	}
}

func (c CreateTransferRequestBodyLogTransferInfoLogTransferMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTransferRequestBodyLogTransferInfoLogTransferMode) UnmarshalJSON(b []byte) error {
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

type CreateTransferRequestBodyLogTransferInfoLogStorageFormat struct {
	value string
}

type CreateTransferRequestBodyLogTransferInfoLogStorageFormatEnum struct {
	JSON CreateTransferRequestBodyLogTransferInfoLogStorageFormat
	RAW  CreateTransferRequestBodyLogTransferInfoLogStorageFormat
}

func GetCreateTransferRequestBodyLogTransferInfoLogStorageFormatEnum() CreateTransferRequestBodyLogTransferInfoLogStorageFormatEnum {
	return CreateTransferRequestBodyLogTransferInfoLogStorageFormatEnum{
		JSON: CreateTransferRequestBodyLogTransferInfoLogStorageFormat{
			value: "JSON",
		},
		RAW: CreateTransferRequestBodyLogTransferInfoLogStorageFormat{
			value: "RAW",
		},
	}
}

func (c CreateTransferRequestBodyLogTransferInfoLogStorageFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTransferRequestBodyLogTransferInfoLogStorageFormat) UnmarshalJSON(b []byte) error {
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

type CreateTransferRequestBodyLogTransferInfoLogTransferStatus struct {
	value string
}

type CreateTransferRequestBodyLogTransferInfoLogTransferStatusEnum struct {
	ENABLE    CreateTransferRequestBodyLogTransferInfoLogTransferStatus
	DISABLE   CreateTransferRequestBodyLogTransferInfoLogTransferStatus
	EXCEPTION CreateTransferRequestBodyLogTransferInfoLogTransferStatus
}

func GetCreateTransferRequestBodyLogTransferInfoLogTransferStatusEnum() CreateTransferRequestBodyLogTransferInfoLogTransferStatusEnum {
	return CreateTransferRequestBodyLogTransferInfoLogTransferStatusEnum{
		ENABLE: CreateTransferRequestBodyLogTransferInfoLogTransferStatus{
			value: "ENABLE",
		},
		DISABLE: CreateTransferRequestBodyLogTransferInfoLogTransferStatus{
			value: "DISABLE",
		},
		EXCEPTION: CreateTransferRequestBodyLogTransferInfoLogTransferStatus{
			value: "EXCEPTION",
		},
	}
}

func (c CreateTransferRequestBodyLogTransferInfoLogTransferStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTransferRequestBodyLogTransferInfoLogTransferStatus) UnmarshalJSON(b []byte) error {
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
