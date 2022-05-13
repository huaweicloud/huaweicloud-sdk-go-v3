package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ResetConsumeOffsetRequest struct {

	// 引擎类型：reliability。
	Engine ResetConsumeOffsetRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组名称。
	GroupId string `json:"group_id"`

	Body *ResetConsumeOffsetReq `json:"body,omitempty"`
}

func (o ResetConsumeOffsetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetRequest struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetRequest", string(data)}, " ")
}

type ResetConsumeOffsetRequestEngine struct {
	value string
}

type ResetConsumeOffsetRequestEngineEnum struct {
	RELIABILITY ResetConsumeOffsetRequestEngine
}

func GetResetConsumeOffsetRequestEngineEnum() ResetConsumeOffsetRequestEngineEnum {
	return ResetConsumeOffsetRequestEngineEnum{
		RELIABILITY: ResetConsumeOffsetRequestEngine{
			value: "reliability",
		},
	}
}

func (c ResetConsumeOffsetRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResetConsumeOffsetRequestEngine) UnmarshalJSON(b []byte) error {
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
