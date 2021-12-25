package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AutoScalingPolicyReqV11 struct {
	// 弹性伸缩规则适用的节点类型，当前只支持task节点。

	NodeGroup AutoScalingPolicyReqV11NodeGroup `json:"node_group"`

	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy"`
}

func (o AutoScalingPolicyReqV11) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoScalingPolicyReqV11 struct{}"
	}

	return strings.Join([]string{"AutoScalingPolicyReqV11", string(data)}, " ")
}

type AutoScalingPolicyReqV11NodeGroup struct {
	value string
}

type AutoScalingPolicyReqV11NodeGroupEnum struct {
	TASK_NODE_DEFAULT_GROUP AutoScalingPolicyReqV11NodeGroup
}

func GetAutoScalingPolicyReqV11NodeGroupEnum() AutoScalingPolicyReqV11NodeGroupEnum {
	return AutoScalingPolicyReqV11NodeGroupEnum{
		TASK_NODE_DEFAULT_GROUP: AutoScalingPolicyReqV11NodeGroup{
			value: "task_node_default_group",
		},
	}
}

func (c AutoScalingPolicyReqV11NodeGroup) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoScalingPolicyReqV11NodeGroup) UnmarshalJSON(b []byte) error {
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
