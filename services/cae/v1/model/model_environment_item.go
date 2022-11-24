package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EnvironmentItem struct {

	// 环境id。
	Id *string `json:"id,omitempty"`

	// 环境名称。
	Name *string `json:"name,omitempty"`

	// 任务id。
	JobId *string `json:"job_id,omitempty"`

	// 环境状态。
	Status *EnvironmentItemStatus `json:"status,omitempty"`

	// 环境类型。
	Type *EnvironmentItemType `json:"type,omitempty"`

	// 环境信息。
	Annotations map[string]string `json:"annotations,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 修改时间。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o EnvironmentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentItem struct{}"
	}

	return strings.Join([]string{"EnvironmentItem", string(data)}, " ")
}

type EnvironmentItemStatus struct {
	value string
}

type EnvironmentItemStatusEnum struct {
	CREATING      EnvironmentItemStatus
	FINISH        EnvironmentItemStatus
	DELETING      EnvironmentItemStatus
	ERROR         EnvironmentItemStatus
	FREEZE        EnvironmentItemStatus
	POLICE_FREEZE EnvironmentItemStatus
	DELETE_FAILED EnvironmentItemStatus
}

func GetEnvironmentItemStatusEnum() EnvironmentItemStatusEnum {
	return EnvironmentItemStatusEnum{
		CREATING: EnvironmentItemStatus{
			value: "creating",
		},
		FINISH: EnvironmentItemStatus{
			value: "finish",
		},
		DELETING: EnvironmentItemStatus{
			value: "deleting",
		},
		ERROR: EnvironmentItemStatus{
			value: "error",
		},
		FREEZE: EnvironmentItemStatus{
			value: "freeze",
		},
		POLICE_FREEZE: EnvironmentItemStatus{
			value: "police_freeze",
		},
		DELETE_FAILED: EnvironmentItemStatus{
			value: "delete_failed",
		},
	}
}

func (c EnvironmentItemStatus) Value() string {
	return c.value
}

func (c EnvironmentItemStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvironmentItemStatus) UnmarshalJSON(b []byte) error {
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

type EnvironmentItemType struct {
	value string
}

type EnvironmentItemTypeEnum struct {
	EXCLUSIVE EnvironmentItemType
}

func GetEnvironmentItemTypeEnum() EnvironmentItemTypeEnum {
	return EnvironmentItemTypeEnum{
		EXCLUSIVE: EnvironmentItemType{
			value: "exclusive",
		},
	}
}

func (c EnvironmentItemType) Value() string {
	return c.value
}

func (c EnvironmentItemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvironmentItemType) UnmarshalJSON(b []byte) error {
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
