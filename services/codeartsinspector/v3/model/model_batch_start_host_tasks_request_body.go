package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchStartHostTasksRequestBody struct {

	// 主机ID列表
	Hosts *[]string `json:"hosts,omitempty"`

	// 对扫描任务的操作:   * start - 启动   * cancel - 取消
	Action *BatchStartHostTasksRequestBodyAction `json:"action,omitempty"`
}

func (o BatchStartHostTasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartHostTasksRequestBody struct{}"
	}

	return strings.Join([]string{"BatchStartHostTasksRequestBody", string(data)}, " ")
}

type BatchStartHostTasksRequestBodyAction struct {
	value string
}

type BatchStartHostTasksRequestBodyActionEnum struct {
	START  BatchStartHostTasksRequestBodyAction
	CANCEL BatchStartHostTasksRequestBodyAction
}

func GetBatchStartHostTasksRequestBodyActionEnum() BatchStartHostTasksRequestBodyActionEnum {
	return BatchStartHostTasksRequestBodyActionEnum{
		START: BatchStartHostTasksRequestBodyAction{
			value: "start",
		},
		CANCEL: BatchStartHostTasksRequestBodyAction{
			value: "cancel",
		},
	}
}

func (c BatchStartHostTasksRequestBodyAction) Value() string {
	return c.value
}

func (c BatchStartHostTasksRequestBodyAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchStartHostTasksRequestBodyAction) UnmarshalJSON(b []byte) error {
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
