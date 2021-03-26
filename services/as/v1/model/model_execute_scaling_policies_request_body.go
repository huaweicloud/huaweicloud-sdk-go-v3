package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 批量操作弹性伸缩策略
type ExecuteScalingPoliciesRequestBody struct {
	// 伸缩策略ID。

	ScalingPolicyId []string `json:"scaling_policy_id"`
	// 是否强制删除伸缩策略。默认为no，可选值为yes或no。只有action为delete时，该字段才生效。

	ForceDelete *string `json:"force_delete,omitempty"`
	// 批量操作伸缩策略action标识：删除：delete。启用：resume。停止：pause。

	Action ExecuteScalingPoliciesRequestBodyAction `json:"action"`
}

func (o ExecuteScalingPoliciesRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExecuteScalingPoliciesRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPoliciesRequestBody", string(data)}, " ")
}

type ExecuteScalingPoliciesRequestBodyAction struct {
	value string
}

type ExecuteScalingPoliciesRequestBodyActionEnum struct {
	DELETE ExecuteScalingPoliciesRequestBodyAction
	RESUME ExecuteScalingPoliciesRequestBodyAction
	PAUSE  ExecuteScalingPoliciesRequestBodyAction
}

func GetExecuteScalingPoliciesRequestBodyActionEnum() ExecuteScalingPoliciesRequestBodyActionEnum {
	return ExecuteScalingPoliciesRequestBodyActionEnum{
		DELETE: ExecuteScalingPoliciesRequestBodyAction{
			value: "delete",
		},
		RESUME: ExecuteScalingPoliciesRequestBodyAction{
			value: "resume",
		},
		PAUSE: ExecuteScalingPoliciesRequestBodyAction{
			value: "pause",
		},
	}
}

func (c ExecuteScalingPoliciesRequestBodyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ExecuteScalingPoliciesRequestBodyAction) UnmarshalJSON(b []byte) error {
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
