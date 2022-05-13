package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchDeleteInstanceReq struct {

	// 实例的ID列表。
	Instances *[]string `json:"instances,omitempty"`

	// 对实例的操作：delete
	Action BatchDeleteInstanceReqAction `json:"action"`

	// 参数值为RocketMQ，表示删除租户所有创建失败的RocketMQ实例。
	AllFailure *BatchDeleteInstanceReqAllFailure `json:"all_failure,omitempty"`
}

func (o BatchDeleteInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceReq", string(data)}, " ")
}

type BatchDeleteInstanceReqAction struct {
	value string
}

type BatchDeleteInstanceReqActionEnum struct {
	DELETE BatchDeleteInstanceReqAction
}

func GetBatchDeleteInstanceReqActionEnum() BatchDeleteInstanceReqActionEnum {
	return BatchDeleteInstanceReqActionEnum{
		DELETE: BatchDeleteInstanceReqAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteInstanceReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteInstanceReqAction) UnmarshalJSON(b []byte) error {
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

type BatchDeleteInstanceReqAllFailure struct {
	value string
}

type BatchDeleteInstanceReqAllFailureEnum struct {
	TRUE  BatchDeleteInstanceReqAllFailure
	FALSE BatchDeleteInstanceReqAllFailure
}

func GetBatchDeleteInstanceReqAllFailureEnum() BatchDeleteInstanceReqAllFailureEnum {
	return BatchDeleteInstanceReqAllFailureEnum{
		TRUE: BatchDeleteInstanceReqAllFailure{
			value: "true",
		},
		FALSE: BatchDeleteInstanceReqAllFailure{
			value: "false",
		},
	}
}

func (c BatchDeleteInstanceReqAllFailure) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteInstanceReqAllFailure) UnmarshalJSON(b []byte) error {
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
