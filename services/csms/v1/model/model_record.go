package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Record 事件通知记录。
type Record struct {

	// 凭据名称。
	EventName *string `json:"event_name,omitempty"`

	// 事件类型。 - SECRET_VERSION_CREATED:版本创建 - SECRET_VERSION_EXPIRED:版本过期 - SECRET_ROTATED:凭据轮转成功 - SECRET_DELETED:凭据删除 - SECRET_ROTATED_FAILED:凭据轮转失败
	TriggerEventType *RecordTriggerEventType `json:"trigger_event_type,omitempty"`

	// 事件通知记录的创建时间，时间戳，即从1970年1月1日至该时间的总秒数。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 凭据名称。
	SecretName *string `json:"secret_name,omitempty"`

	// 凭据类型  - COMMON：通用凭据(默认)。用于应用系统中的各种敏感信息储存。 - RDS：RDS凭据 。专门针对RDS的凭据，用于存储RDS的账号信息。（已不支持，使用RDS-FG替代） - RDS-FG：RDS凭据 。专门针对RDS的凭据，用于存储RDS的账号信息。 - GaussDB-FG：GaussDB凭据。专门针对GaussDB的凭据，用于存储GaussDB的账号信息。
	SecretType *RecordSecretType `json:"secret_type,omitempty"`

	// 事件通知的对象名称。
	NotificationTargetName *string `json:"notification_target_name,omitempty"`

	// 事件通知的对象ID。
	NotificationTargetId *string `json:"notification_target_id,omitempty"`

	// 事件通知的内容。
	NotificationContent *string `json:"notification_content,omitempty"`

	// 事件通知状态。  - SUCCESS：事件通知成功。 - FAIL：事件通知失败。 - INVALID：事件通知配置主题信息无效或不正确，无法触发通知。
	NotificationStatus *RecordNotificationStatus `json:"notification_status,omitempty"`
}

func (o Record) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Record struct{}"
	}

	return strings.Join([]string{"Record", string(data)}, " ")
}

type RecordTriggerEventType struct {
	value string
}

type RecordTriggerEventTypeEnum struct {
	SECRET_VERSION_CREATED RecordTriggerEventType
	SECRET_VERSION_EXPIRED RecordTriggerEventType
	SECRET_ROTATED         RecordTriggerEventType
	SECRET_DELETED         RecordTriggerEventType
	SECRET_ROTATED_FAILED  RecordTriggerEventType
}

func GetRecordTriggerEventTypeEnum() RecordTriggerEventTypeEnum {
	return RecordTriggerEventTypeEnum{
		SECRET_VERSION_CREATED: RecordTriggerEventType{
			value: "SECRET_VERSION_CREATED",
		},
		SECRET_VERSION_EXPIRED: RecordTriggerEventType{
			value: "SECRET_VERSION_EXPIRED",
		},
		SECRET_ROTATED: RecordTriggerEventType{
			value: "SECRET_ROTATED",
		},
		SECRET_DELETED: RecordTriggerEventType{
			value: "SECRET_DELETED",
		},
		SECRET_ROTATED_FAILED: RecordTriggerEventType{
			value: "SECRET_ROTATED_FAILED",
		},
	}
}

func (c RecordTriggerEventType) Value() string {
	return c.value
}

func (c RecordTriggerEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordTriggerEventType) UnmarshalJSON(b []byte) error {
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

type RecordSecretType struct {
	value string
}

type RecordSecretTypeEnum struct {
	COMMON      RecordSecretType
	RDS_FG      RecordSecretType
	GAUSS_DB_FG RecordSecretType
}

func GetRecordSecretTypeEnum() RecordSecretTypeEnum {
	return RecordSecretTypeEnum{
		COMMON: RecordSecretType{
			value: "COMMON",
		},
		RDS_FG: RecordSecretType{
			value: "RDS-FG",
		},
		GAUSS_DB_FG: RecordSecretType{
			value: "GaussDB-FG",
		},
	}
}

func (c RecordSecretType) Value() string {
	return c.value
}

func (c RecordSecretType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordSecretType) UnmarshalJSON(b []byte) error {
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

type RecordNotificationStatus struct {
	value string
}

type RecordNotificationStatusEnum struct {
	SUCCESS RecordNotificationStatus
	FAIL    RecordNotificationStatus
	INVALID RecordNotificationStatus
}

func GetRecordNotificationStatusEnum() RecordNotificationStatusEnum {
	return RecordNotificationStatusEnum{
		SUCCESS: RecordNotificationStatus{
			value: "SUCCESS",
		},
		FAIL: RecordNotificationStatus{
			value: "FAIL",
		},
		INVALID: RecordNotificationStatus{
			value: "INVALID",
		},
	}
}

func (c RecordNotificationStatus) Value() string {
	return c.value
}

func (c RecordNotificationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordNotificationStatus) UnmarshalJSON(b []byte) error {
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
