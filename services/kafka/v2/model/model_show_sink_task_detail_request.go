package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSinkTaskDetailRequest Request Object
type ShowSinkTaskDetailRequest struct {

	// 实例转储ID。  请参考[查询实例](ShowInstance.xml)返回的数据。
	ConnectorId string `json:"connector_id"`

	// 转储任务ID。
	TaskId string `json:"task_id"`

	// 是否包含topic信息。默认是false。
	TopicInfo *ShowSinkTaskDetailRequestTopicInfo `json:"topic-info,omitempty"`
}

func (o ShowSinkTaskDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSinkTaskDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSinkTaskDetailRequest", string(data)}, " ")
}

type ShowSinkTaskDetailRequestTopicInfo struct {
	value string
}

type ShowSinkTaskDetailRequestTopicInfoEnum struct {
	TRUE  ShowSinkTaskDetailRequestTopicInfo
	FALSE ShowSinkTaskDetailRequestTopicInfo
}

func GetShowSinkTaskDetailRequestTopicInfoEnum() ShowSinkTaskDetailRequestTopicInfoEnum {
	return ShowSinkTaskDetailRequestTopicInfoEnum{
		TRUE: ShowSinkTaskDetailRequestTopicInfo{
			value: "true",
		},
		FALSE: ShowSinkTaskDetailRequestTopicInfo{
			value: "false",
		},
	}
}

func (c ShowSinkTaskDetailRequestTopicInfo) Value() string {
	return c.value
}

func (c ShowSinkTaskDetailRequestTopicInfo) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSinkTaskDetailRequestTopicInfo) UnmarshalJSON(b []byte) error {
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
