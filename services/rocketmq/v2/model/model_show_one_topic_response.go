package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowOneTopicResponse Response Object
type ShowOneTopicResponse struct {

	// topic名称。
	Name *string `json:"name,omitempty"`

	// 总读队列个数。
	TotalReadQueueNum float32 `json:"total_read_queue_num,omitempty"`

	// 总写队列个数。
	TotalWriteQueueNum float32 `json:"total_write_queue_num,omitempty"`

	// 权限。
	Permission *ShowOneTopicResponsePermission `json:"permission,omitempty"`

	// 关联的代理。
	Brokers        *[]TopicBrokers `json:"brokers,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowOneTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOneTopicResponse struct{}"
	}

	return strings.Join([]string{"ShowOneTopicResponse", string(data)}, " ")
}

type ShowOneTopicResponsePermission struct {
	value string
}

type ShowOneTopicResponsePermissionEnum struct {
	SUB ShowOneTopicResponsePermission
	PUB ShowOneTopicResponsePermission
	ALL ShowOneTopicResponsePermission
}

func GetShowOneTopicResponsePermissionEnum() ShowOneTopicResponsePermissionEnum {
	return ShowOneTopicResponsePermissionEnum{
		SUB: ShowOneTopicResponsePermission{
			value: "sub",
		},
		PUB: ShowOneTopicResponsePermission{
			value: "pub",
		},
		ALL: ShowOneTopicResponsePermission{
			value: "all",
		},
	}
}

func (c ShowOneTopicResponsePermission) Value() string {
	return c.value
}

func (c ShowOneTopicResponsePermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowOneTopicResponsePermission) UnmarshalJSON(b []byte) error {
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
