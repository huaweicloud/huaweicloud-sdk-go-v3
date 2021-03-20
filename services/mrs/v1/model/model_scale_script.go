package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScaleScript struct {
	// 弹性伸缩自定义自动化脚本的名称，同一个集群的自定义自动化脚本名称不允许相同。  只能由数字、英文字符、空格、中划线和下划线组成，且不能以空格开头。  可输入的字符串长度为1～64个字符。

	Name string `json:"name"`
	// 自定义自动化脚本的路径。设置为OBS桶的路径或虚拟机本地的路径。  OBS桶的路径：直接手动输入脚本路径。示例：s3a://XXX/scale.sh 虚拟机本地的路径：用户需要输入正确的脚本路径。脚本所在的路径必须以‘/’开头，以.sh结尾。

	Uri string `json:"uri"`
	// 自定义自动化脚本参数。  多个参数间用空格隔开。 可以传入以下系统预定义参数： ${mrs_scale_node_num}：扩缩容节点数 ${mrs_scale_type}：扩缩容类型，扩容为scale_out，缩容为scale_in ${mrs_scale_node_hostnames}：扩缩容的节点主机名称 ${mrs_scale_node_ips}：扩缩容的节点IP ${mrs_scale_rule_name}：触发扩缩容的规则名 其他用户自定义参数使用方式与普通shell脚本相同，多个参数中间用空格隔开。

	Parameters *string `json:"parameters,omitempty"`
	// 自定义自动化脚本所执行的节点类型，包含Master、Core和Task三种类型。

	Nodes []string `json:"nodes"`
	// 自定义自动化脚本是否只运行在主Master节点上。  缺省值为false，表示自定义自动化脚本可运行在所有Master节点上。

	ActiveMaster *bool `json:"active_master,omitempty"`
	// 自自定义自动化脚本执行失败后，是否继续执行后续脚本和创建集群。  continue：继续执行后续脚本。 errorout：终止操作。 说明： 建议您在调试阶段设置为“continue”，无论此自定义自动化脚本是否执行成功，则集群都能继续安装和启动。 由于缩容成功无法回滚，因此缩容后执行的脚本“fail_action”必须设置为“continue”。

	FailAction ScaleScriptFailAction `json:"fail_action"`
	// 脚本执行时机。  支持以下四个阶段：  before_scale_out：扩容前 before_scale_in：缩容前 after_scale_out：扩容后 after_scale_in：缩容后

	ActionStage ScaleScriptActionStage `json:"action_stage"`
}

func (o ScaleScript) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ScaleScript struct{}"
	}

	return strings.Join([]string{"ScaleScript", string(data)}, " ")
}

type ScaleScriptFailAction struct {
	value string
}

type ScaleScriptFailActionEnum struct {
	CONTINUE ScaleScriptFailAction
	ERROROUT ScaleScriptFailAction
}

func GetScaleScriptFailActionEnum() ScaleScriptFailActionEnum {
	return ScaleScriptFailActionEnum{
		CONTINUE: ScaleScriptFailAction{
			value: "continue",
		},
		ERROROUT: ScaleScriptFailAction{
			value: "errorout",
		},
	}
}

func (c ScaleScriptFailAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ScaleScriptFailAction) UnmarshalJSON(b []byte) error {
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

type ScaleScriptActionStage struct {
	value string
}

type ScaleScriptActionStageEnum struct {
	BEFORE_SCALE_OUT ScaleScriptActionStage
	BEFORE_SCALE_IN  ScaleScriptActionStage
	AFTER_SCALE_OUT  ScaleScriptActionStage
	AFTER_SCALE_IN   ScaleScriptActionStage
}

func GetScaleScriptActionStageEnum() ScaleScriptActionStageEnum {
	return ScaleScriptActionStageEnum{
		BEFORE_SCALE_OUT: ScaleScriptActionStage{
			value: "before_scale_out",
		},
		BEFORE_SCALE_IN: ScaleScriptActionStage{
			value: "before_scale_in",
		},
		AFTER_SCALE_OUT: ScaleScriptActionStage{
			value: "after_scale_out",
		},
		AFTER_SCALE_IN: ScaleScriptActionStage{
			value: "after_scale_in",
		},
	}
}

func (c ScaleScriptActionStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ScaleScriptActionStage) UnmarshalJSON(b []byte) error {
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
