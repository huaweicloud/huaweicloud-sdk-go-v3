package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UnblockRecordResponseUnblockRecord struct {

	// ip地址
	Ip *string `json:"ip,omitempty"`

	// 执行者
	Executor *string `json:"executor,omitempty"`

	// 封堵id
	BlockId *int64 `json:"block_id,omitempty"`

	// 封堵时间
	BlockingTime *int64 `json:"blocking_time,omitempty"`

	// 解封时间
	UnblockingTime *int64 `json:"unblocking_time,omitempty"`

	// 解封类型。manual：人工；automatic：自动
	UnblockType *UnblockRecordResponseUnblockRecordUnblockType `json:"unblock_type,omitempty"`

	// 状态。unblocking：解封中；success：成功；failed：失败
	Status *UnblockRecordResponseUnblockRecordStatus `json:"status,omitempty"`

	// 时间
	SortTime *int64 `json:"sort_time,omitempty"`
}

func (o UnblockRecordResponseUnblockRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnblockRecordResponseUnblockRecord struct{}"
	}

	return strings.Join([]string{"UnblockRecordResponseUnblockRecord", string(data)}, " ")
}

type UnblockRecordResponseUnblockRecordUnblockType struct {
	value string
}

type UnblockRecordResponseUnblockRecordUnblockTypeEnum struct {
	MANUAL    UnblockRecordResponseUnblockRecordUnblockType
	AUTOMATIC UnblockRecordResponseUnblockRecordUnblockType
}

func GetUnblockRecordResponseUnblockRecordUnblockTypeEnum() UnblockRecordResponseUnblockRecordUnblockTypeEnum {
	return UnblockRecordResponseUnblockRecordUnblockTypeEnum{
		MANUAL: UnblockRecordResponseUnblockRecordUnblockType{
			value: "manual",
		},
		AUTOMATIC: UnblockRecordResponseUnblockRecordUnblockType{
			value: "automatic",
		},
	}
}

func (c UnblockRecordResponseUnblockRecordUnblockType) Value() string {
	return c.value
}

func (c UnblockRecordResponseUnblockRecordUnblockType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UnblockRecordResponseUnblockRecordUnblockType) UnmarshalJSON(b []byte) error {
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

type UnblockRecordResponseUnblockRecordStatus struct {
	value string
}

type UnblockRecordResponseUnblockRecordStatusEnum struct {
	UNBLOCKING UnblockRecordResponseUnblockRecordStatus
	SUCCESS    UnblockRecordResponseUnblockRecordStatus
	FAILED     UnblockRecordResponseUnblockRecordStatus
}

func GetUnblockRecordResponseUnblockRecordStatusEnum() UnblockRecordResponseUnblockRecordStatusEnum {
	return UnblockRecordResponseUnblockRecordStatusEnum{
		UNBLOCKING: UnblockRecordResponseUnblockRecordStatus{
			value: "unblocking",
		},
		SUCCESS: UnblockRecordResponseUnblockRecordStatus{
			value: "success",
		},
		FAILED: UnblockRecordResponseUnblockRecordStatus{
			value: "failed",
		},
	}
}

func (c UnblockRecordResponseUnblockRecordStatus) Value() string {
	return c.value
}

func (c UnblockRecordResponseUnblockRecordStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UnblockRecordResponseUnblockRecordStatus) UnmarshalJSON(b []byte) error {
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
