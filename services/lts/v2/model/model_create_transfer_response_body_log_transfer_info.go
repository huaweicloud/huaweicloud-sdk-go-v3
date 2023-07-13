package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateTransferResponseBodyLogTransferInfo 日志转储信息
type CreateTransferResponseBodyLogTransferInfo struct {
	LogAgencyTransfer *CreateTransferResponseBodyLogTransferInfoLogAgencyTransfer `json:"log_agency_transfer,omitempty"`

	// 日志转储创建时间
	LogCreateTime int64 `json:"log_create_time"`

	// 日志转储格式。只支持\"RAW\", \"JSON\"。RAW是指原始日志格式，JSON是指JSON日志格式。OBS转储和DIS转储支持JSON和RAW，DMS转储仅支持RAW
	LogStorageFormat CreateTransferResponseBodyLogTransferInfoLogStorageFormat `json:"log_storage_format"`

	// 日志转储详细信息
	LogTransferDetail *TransferDetail `json:"log_transfer_detail"`

	// 日志转储方式。cycle是指周期性转储，realTime是指实时转储。OBS转储只支持\"cycle\"，DIS转储和DMS转储只支持\"realTime\"。
	LogTransferMode CreateTransferResponseBodyLogTransferInfoLogTransferMode `json:"log_transfer_mode"`

	// 日志转储状态，ENABLE是指日志转储开启状态，DISABLE是指日志转储关闭状态，EXCEPTION是指日志转储异常状态
	LogTransferStatus CreateTransferResponseBodyLogTransferInfoLogTransferStatus `json:"log_transfer_status"`

	// 日志转储类型。OBS指OBS日志转储，DIS指DIS日志转储，DMS指DMS日志转储
	LogTransferType CreateTransferResponseBodyLogTransferInfoLogTransferType `json:"log_transfer_type"`
}

func (o CreateTransferResponseBodyLogTransferInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransferResponseBodyLogTransferInfo struct{}"
	}

	return strings.Join([]string{"CreateTransferResponseBodyLogTransferInfo", string(data)}, " ")
}

type CreateTransferResponseBodyLogTransferInfoLogStorageFormat struct {
	value string
}

type CreateTransferResponseBodyLogTransferInfoLogStorageFormatEnum struct {
	JSON CreateTransferResponseBodyLogTransferInfoLogStorageFormat
	RAW  CreateTransferResponseBodyLogTransferInfoLogStorageFormat
}

func GetCreateTransferResponseBodyLogTransferInfoLogStorageFormatEnum() CreateTransferResponseBodyLogTransferInfoLogStorageFormatEnum {
	return CreateTransferResponseBodyLogTransferInfoLogStorageFormatEnum{
		JSON: CreateTransferResponseBodyLogTransferInfoLogStorageFormat{
			value: "JSON",
		},
		RAW: CreateTransferResponseBodyLogTransferInfoLogStorageFormat{
			value: "RAW",
		},
	}
}

func (c CreateTransferResponseBodyLogTransferInfoLogStorageFormat) Value() string {
	return c.value
}

func (c CreateTransferResponseBodyLogTransferInfoLogStorageFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTransferResponseBodyLogTransferInfoLogStorageFormat) UnmarshalJSON(b []byte) error {
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

type CreateTransferResponseBodyLogTransferInfoLogTransferMode struct {
	value string
}

type CreateTransferResponseBodyLogTransferInfoLogTransferModeEnum struct {
	CYCLE     CreateTransferResponseBodyLogTransferInfoLogTransferMode
	REAL_TIME CreateTransferResponseBodyLogTransferInfoLogTransferMode
}

func GetCreateTransferResponseBodyLogTransferInfoLogTransferModeEnum() CreateTransferResponseBodyLogTransferInfoLogTransferModeEnum {
	return CreateTransferResponseBodyLogTransferInfoLogTransferModeEnum{
		CYCLE: CreateTransferResponseBodyLogTransferInfoLogTransferMode{
			value: "cycle",
		},
		REAL_TIME: CreateTransferResponseBodyLogTransferInfoLogTransferMode{
			value: "realTime",
		},
	}
}

func (c CreateTransferResponseBodyLogTransferInfoLogTransferMode) Value() string {
	return c.value
}

func (c CreateTransferResponseBodyLogTransferInfoLogTransferMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTransferResponseBodyLogTransferInfoLogTransferMode) UnmarshalJSON(b []byte) error {
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

type CreateTransferResponseBodyLogTransferInfoLogTransferStatus struct {
	value string
}

type CreateTransferResponseBodyLogTransferInfoLogTransferStatusEnum struct {
	ENABLE    CreateTransferResponseBodyLogTransferInfoLogTransferStatus
	DISABLE   CreateTransferResponseBodyLogTransferInfoLogTransferStatus
	EXCEPTION CreateTransferResponseBodyLogTransferInfoLogTransferStatus
}

func GetCreateTransferResponseBodyLogTransferInfoLogTransferStatusEnum() CreateTransferResponseBodyLogTransferInfoLogTransferStatusEnum {
	return CreateTransferResponseBodyLogTransferInfoLogTransferStatusEnum{
		ENABLE: CreateTransferResponseBodyLogTransferInfoLogTransferStatus{
			value: "ENABLE",
		},
		DISABLE: CreateTransferResponseBodyLogTransferInfoLogTransferStatus{
			value: "DISABLE",
		},
		EXCEPTION: CreateTransferResponseBodyLogTransferInfoLogTransferStatus{
			value: "EXCEPTION",
		},
	}
}

func (c CreateTransferResponseBodyLogTransferInfoLogTransferStatus) Value() string {
	return c.value
}

func (c CreateTransferResponseBodyLogTransferInfoLogTransferStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTransferResponseBodyLogTransferInfoLogTransferStatus) UnmarshalJSON(b []byte) error {
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

type CreateTransferResponseBodyLogTransferInfoLogTransferType struct {
	value string
}

type CreateTransferResponseBodyLogTransferInfoLogTransferTypeEnum struct {
	OBS CreateTransferResponseBodyLogTransferInfoLogTransferType
	DIS CreateTransferResponseBodyLogTransferInfoLogTransferType
	DMS CreateTransferResponseBodyLogTransferInfoLogTransferType
}

func GetCreateTransferResponseBodyLogTransferInfoLogTransferTypeEnum() CreateTransferResponseBodyLogTransferInfoLogTransferTypeEnum {
	return CreateTransferResponseBodyLogTransferInfoLogTransferTypeEnum{
		OBS: CreateTransferResponseBodyLogTransferInfoLogTransferType{
			value: "OBS",
		},
		DIS: CreateTransferResponseBodyLogTransferInfoLogTransferType{
			value: "DIS",
		},
		DMS: CreateTransferResponseBodyLogTransferInfoLogTransferType{
			value: "DMS",
		},
	}
}

func (c CreateTransferResponseBodyLogTransferInfoLogTransferType) Value() string {
	return c.value
}

func (c CreateTransferResponseBodyLogTransferInfoLogTransferType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTransferResponseBodyLogTransferInfoLogTransferType) UnmarshalJSON(b []byte) error {
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
