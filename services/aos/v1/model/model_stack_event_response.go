package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// stack_event
type StackEventResponse struct {

	// 资源的id名称，即对应资源作为唯一id使用的值的名称，当资源未创建的时候，不返回resource_id_key
	ResourceIdKey *string `json:"resource_id_key,omitempty"`

	// 资源的id值，即对应资源作为唯一id使用的值，当资源未创建的时候，不返回resource_id_value
	ResourceIdValue *string `json:"resource_id_value,omitempty"`

	// 资源的名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源的类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 事件发生的时间，格式遵循RFC3339，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z
	Time *string `json:"time,omitempty"`

	// 此次事件的类型 * `LOG` - 记录状态信息，比如当前状态，目标状态等。详细信息将在EventMessage中存储。此为暂时状态， 因为目前我们无法解析terraform的log，只能直全部以LOG形式打出，未来我们吃掉terraform内核以后， 将去除这个状态，并改为下列的更精准的状态 * `ERROR` - 记录失败信息 * `DRIFT` - 记录资源偏移信息 * `SUMMARY` - 记录资源变更结果总结 * `CREATION_IN_PROGRESS` - 正在生成 * `CREATION_FAILED` - 生成失败 * `CREATION_COMPLETE` - 生成完成 * `DELETION_IN_PROGRESS` - 正在删除 * `DELETION_FAILED` - 删除失败 * `DELETION_COMPLETE` - 已经删除 * `DELETION_SKIPPED` - 跳过删除。未来我们将支持，用户可以从IaC中删除，但是不真的删除资源本身 * `UPDATE_IN_PROGRESS` - 正在更新。此处的更新特指非替换式更新，如果是替换式更新，则使用CREATION后DELETION * `UPDATE_FAILED` - 更新失败。此处的更新特指非替换式更新，如果是替换式更新，则使用CREATION后 * `UPDATE_COMPLETE` - 更新完成。此处的更新特指非替换式更新，如果是替换式更新，则使用CREATION后DELETION
	EventType *StackEventResponseEventType `json:"event_type,omitempty"`

	// 事件对应的详细信息
	EventMessage *string `json:"event_message,omitempty"`

	// 资源改动所花的时间，以秒为单位
	ElapsedSeconds *int32 `json:"elapsed_seconds,omitempty"`
}

func (o StackEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackEventResponse struct{}"
	}

	return strings.Join([]string{"StackEventResponse", string(data)}, " ")
}

type StackEventResponseEventType struct {
	value string
}

type StackEventResponseEventTypeEnum struct {
	LOG                  StackEventResponseEventType
	ERROR                StackEventResponseEventType
	DRIFT                StackEventResponseEventType
	SUMMARY              StackEventResponseEventType
	CREATION_IN_PROGRESS StackEventResponseEventType
	CREATION_FAILED      StackEventResponseEventType
	CREATION_COMPLETE    StackEventResponseEventType
	DELETION_IN_PROGRESS StackEventResponseEventType
	DELETION_FAILED      StackEventResponseEventType
	DELETION_COMPLETE    StackEventResponseEventType
	DELETION_SKIPPED     StackEventResponseEventType
	UPDATE_IN_PROGRESS   StackEventResponseEventType
	UPDATE_FAILED        StackEventResponseEventType
	UPDATE_COMPLETE      StackEventResponseEventType
}

func GetStackEventResponseEventTypeEnum() StackEventResponseEventTypeEnum {
	return StackEventResponseEventTypeEnum{
		LOG: StackEventResponseEventType{
			value: "LOG",
		},
		ERROR: StackEventResponseEventType{
			value: "ERROR",
		},
		DRIFT: StackEventResponseEventType{
			value: "DRIFT",
		},
		SUMMARY: StackEventResponseEventType{
			value: "SUMMARY",
		},
		CREATION_IN_PROGRESS: StackEventResponseEventType{
			value: "CREATION_IN_PROGRESS",
		},
		CREATION_FAILED: StackEventResponseEventType{
			value: "CREATION_FAILED",
		},
		CREATION_COMPLETE: StackEventResponseEventType{
			value: "CREATION_COMPLETE",
		},
		DELETION_IN_PROGRESS: StackEventResponseEventType{
			value: "DELETION_IN_PROGRESS",
		},
		DELETION_FAILED: StackEventResponseEventType{
			value: "DELETION_FAILED",
		},
		DELETION_COMPLETE: StackEventResponseEventType{
			value: "DELETION_COMPLETE",
		},
		DELETION_SKIPPED: StackEventResponseEventType{
			value: "DELETION_SKIPPED",
		},
		UPDATE_IN_PROGRESS: StackEventResponseEventType{
			value: "UPDATE_IN_PROGRESS",
		},
		UPDATE_FAILED: StackEventResponseEventType{
			value: "UPDATE_FAILED",
		},
		UPDATE_COMPLETE: StackEventResponseEventType{
			value: "UPDATE_COMPLETE",
		},
	}
}

func (c StackEventResponseEventType) Value() string {
	return c.value
}

func (c StackEventResponseEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackEventResponseEventType) UnmarshalJSON(b []byte) error {
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
