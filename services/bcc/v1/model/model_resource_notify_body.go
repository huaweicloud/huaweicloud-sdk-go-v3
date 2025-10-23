package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceNotifyBody struct {

	// 事件ID
	EventId string `json:"event_id"`

	// 捕获时间
	CaptureTime *string `json:"capture_time,omitempty"`

	// 账户ID
	DomainId string `json:"domain_id"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 事件类型：CREATE-创建，UPDATE-更新，DELETE-删除
	EventType ResourceNotifyBodyEventType `json:"event_type"`

	// 检验码
	Checksum *string `json:"checksum,omitempty"`

	// 同步类型
	SyncType *string `json:"sync_type,omitempty"`

	Resource *ResourceNotifyEntity `json:"resource,omitempty"`
}

func (o ResourceNotifyBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceNotifyBody struct{}"
	}

	return strings.Join([]string{"ResourceNotifyBody", string(data)}, " ")
}

type ResourceNotifyBodyEventType struct {
	value string
}

type ResourceNotifyBodyEventTypeEnum struct {
	CREATE ResourceNotifyBodyEventType
	UPDATE ResourceNotifyBodyEventType
	DELETE ResourceNotifyBodyEventType
}

func GetResourceNotifyBodyEventTypeEnum() ResourceNotifyBodyEventTypeEnum {
	return ResourceNotifyBodyEventTypeEnum{
		CREATE: ResourceNotifyBodyEventType{
			value: "CREATE",
		},
		UPDATE: ResourceNotifyBodyEventType{
			value: "UPDATE",
		},
		DELETE: ResourceNotifyBodyEventType{
			value: "DELETE",
		},
	}
}

func (c ResourceNotifyBodyEventType) Value() string {
	return c.value
}

func (c ResourceNotifyBodyEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceNotifyBodyEventType) UnmarshalJSON(b []byte) error {
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
