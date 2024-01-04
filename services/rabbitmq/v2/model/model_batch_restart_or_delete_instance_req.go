package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchRestartOrDeleteInstanceReq struct {

	// 实例的ID列表。
	Instances *[]string `json:"instances,omitempty"`

	// 对实例的操作：delete
	Action BatchRestartOrDeleteInstanceReqAction `json:"action"`

	// 是否批量删除创建失败的实例。  当参数值为“rabbitmq”时，删除租户所有创建失败的实例，此时请求参数instances可为空。
	AllFailure *BatchRestartOrDeleteInstanceReqAllFailure `json:"all_failure,omitempty"`
}

func (o BatchRestartOrDeleteInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestartOrDeleteInstanceReq struct{}"
	}

	return strings.Join([]string{"BatchRestartOrDeleteInstanceReq", string(data)}, " ")
}

type BatchRestartOrDeleteInstanceReqAction struct {
	value string
}

type BatchRestartOrDeleteInstanceReqActionEnum struct {
	DELETE BatchRestartOrDeleteInstanceReqAction
}

func GetBatchRestartOrDeleteInstanceReqActionEnum() BatchRestartOrDeleteInstanceReqActionEnum {
	return BatchRestartOrDeleteInstanceReqActionEnum{
		DELETE: BatchRestartOrDeleteInstanceReqAction{
			value: "delete",
		},
	}
}

func (c BatchRestartOrDeleteInstanceReqAction) Value() string {
	return c.value
}

func (c BatchRestartOrDeleteInstanceReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRestartOrDeleteInstanceReqAction) UnmarshalJSON(b []byte) error {
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

type BatchRestartOrDeleteInstanceReqAllFailure struct {
	value string
}

type BatchRestartOrDeleteInstanceReqAllFailureEnum struct {
	RABBITMQ BatchRestartOrDeleteInstanceReqAllFailure
}

func GetBatchRestartOrDeleteInstanceReqAllFailureEnum() BatchRestartOrDeleteInstanceReqAllFailureEnum {
	return BatchRestartOrDeleteInstanceReqAllFailureEnum{
		RABBITMQ: BatchRestartOrDeleteInstanceReqAllFailure{
			value: "rabbitmq",
		},
	}
}

func (c BatchRestartOrDeleteInstanceReqAllFailure) Value() string {
	return c.value
}

func (c BatchRestartOrDeleteInstanceReqAllFailure) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRestartOrDeleteInstanceReqAllFailure) UnmarshalJSON(b []byte) error {
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
