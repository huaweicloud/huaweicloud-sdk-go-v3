package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateTopicReq struct {

	// 总读队列个数。
	ReadQueueNum float32 `json:"read_queue_num,omitempty" xml:"read_queue_num"`

	// 总写队列个数。
	WriteQueueNum float32 `json:"write_queue_num,omitempty" xml:"write_queue_num"`

	// 权限。
	Permission *UpdateTopicReqPermission `json:"permission,omitempty" xml:"permission"`
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
