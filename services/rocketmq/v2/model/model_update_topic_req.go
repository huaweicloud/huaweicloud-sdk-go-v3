package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateTopicReq struct {

	// **参数解释**： 总读队列个数。 **约束限制**： 仅4.8.0实例支持修改该参数。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ReadQueueNum float32 `json:"read_queue_num,omitempty"`

	// **参数解释**： 总写队列个数。 **约束限制**： 仅4.8.0实例支持修改该参数。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	WriteQueueNum float32 `json:"write_queue_num,omitempty"`

	// **参数解释**： 权限。 **约束限制**： 仅4.8.0实例支持修改该参数。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Permission *UpdateTopicReqPermission `json:"permission,omitempty"`

	// **参数解释**： 队列。 **约束限制**： 仅4.8.0实例支持修改该参数。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Queues *[]UpdateTopicQueueEntity `json:"queues,omitempty"`

	// **参数解释**： Topic描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TopicDesc *string `json:"topic_desc,omitempty"`
}

func (o UpdateTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicReq struct{}"
	}

	return strings.Join([]string{"UpdateTopicReq", string(data)}, " ")
}

type UpdateTopicReqPermission struct {
	value string
}

type UpdateTopicReqPermissionEnum struct {
	SUB UpdateTopicReqPermission
	PUB UpdateTopicReqPermission
	ALL UpdateTopicReqPermission
}

func GetUpdateTopicReqPermissionEnum() UpdateTopicReqPermissionEnum {
	return UpdateTopicReqPermissionEnum{
		SUB: UpdateTopicReqPermission{
			value: "sub",
		},
		PUB: UpdateTopicReqPermission{
			value: "pub",
		},
		ALL: UpdateTopicReqPermission{
			value: "all",
		},
	}
}

func (c UpdateTopicReqPermission) Value() string {
	return c.value
}

func (c UpdateTopicReqPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTopicReqPermission) UnmarshalJSON(b []byte) error {
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
