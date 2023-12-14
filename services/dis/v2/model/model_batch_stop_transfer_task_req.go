package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchStopTransferTaskReq struct {

	// 转储任务操作，目前支持：  - stop：停止转储任务
	Action BatchStopTransferTaskReqAction `json:"action"`

	// 待暂停的转储任务列表。
	Tasks []BatchTransferTask `json:"tasks"`
}

func (o BatchStopTransferTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopTransferTaskReq struct{}"
	}

	return strings.Join([]string{"BatchStopTransferTaskReq", string(data)}, " ")
}

type BatchStopTransferTaskReqAction struct {
	value string
}

type BatchStopTransferTaskReqActionEnum struct {
	STOP BatchStopTransferTaskReqAction
}

func GetBatchStopTransferTaskReqActionEnum() BatchStopTransferTaskReqActionEnum {
	return BatchStopTransferTaskReqActionEnum{
		STOP: BatchStopTransferTaskReqAction{
			value: "stop",
		},
	}
}

func (c BatchStopTransferTaskReqAction) Value() string {
	return c.value
}

func (c BatchStopTransferTaskReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchStopTransferTaskReqAction) UnmarshalJSON(b []byte) error {
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
