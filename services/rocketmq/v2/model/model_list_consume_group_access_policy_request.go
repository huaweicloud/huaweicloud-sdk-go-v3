package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListConsumeGroupAccessPolicyRequest Request Object
type ListConsumeGroupAccessPolicyRequest struct {

	// 消息引擎。
	Engine ListConsumeGroupAccessPolicyRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组。
	GroupId string `json:"group_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *string `json:"offset,omitempty"`

	// 查询数量。
	Limit *string `json:"limit,omitempty"`
}

func (o ListConsumeGroupAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumeGroupAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListConsumeGroupAccessPolicyRequest", string(data)}, " ")
}

type ListConsumeGroupAccessPolicyRequestEngine struct {
	value string
}

type ListConsumeGroupAccessPolicyRequestEngineEnum struct {
	RELIABILITY ListConsumeGroupAccessPolicyRequestEngine
}

func GetListConsumeGroupAccessPolicyRequestEngineEnum() ListConsumeGroupAccessPolicyRequestEngineEnum {
	return ListConsumeGroupAccessPolicyRequestEngineEnum{
		RELIABILITY: ListConsumeGroupAccessPolicyRequestEngine{
			value: "reliability",
		},
	}
}

func (c ListConsumeGroupAccessPolicyRequestEngine) Value() string {
	return c.value
}

func (c ListConsumeGroupAccessPolicyRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConsumeGroupAccessPolicyRequestEngine) UnmarshalJSON(b []byte) error {
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
