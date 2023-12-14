package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchStartTransferTaskReq struct {

	// 转储任务操作，目前支持：  - start：启动转储任务
	Action BatchStartTransferTaskReqAction `json:"action"`

	// 待操作的转储任务列表。
	Tasks []BatchTransferTask `json:"tasks"`
}

func (o BatchStartTransferTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartTransferTaskReq struct{}"
	}

	return strings.Join([]string{"BatchStartTransferTaskReq", string(data)}, " ")
}

type BatchStartTransferTaskReqAction struct {
	value string
}

type BatchStartTransferTaskReqActionEnum struct {
	START BatchStartTransferTaskReqAction
}

func GetBatchStartTransferTaskReqActionEnum() BatchStartTransferTaskReqActionEnum {
	return BatchStartTransferTaskReqActionEnum{
		START: BatchStartTransferTaskReqAction{
			value: "start",
		},
	}
}

func (c BatchStartTransferTaskReqAction) Value() string {
	return c.value
}

func (c BatchStartTransferTaskReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchStartTransferTaskReqAction) UnmarshalJSON(b []byte) error {
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
