package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BootstrapScriptResp struct {
	// 引导操作脚本的名称，同一个集群的引导操作脚本名称不允许相同。  只能由数字、英文字符、空格、中划线和下划线组成，且不能以空格开头。  可输入的字符串长度为1～64个字符。

	Name string `json:"name"`
	// 引导操作脚本的路径。设置为OBS桶的路径或虚拟机本地的路径。 - OBS桶的路径：直接手动输入脚本路径。例如输入MRS提供的公共样例脚本路径。示例：s3a://bootstrap/presto/presto-install.sh，其中安装dualroles时，presto-install.sh脚本参数为dualroles, 安装worker时，presto-install.sh脚本参数为worker。根据Presto使用习惯，建议您在Active Master节点上安装dualroles，在Core节点上安装worker。 - 虚拟机本地的路径：用户需要输入正确的脚本路径。脚本所在的路径必须以‘/’开头，以.sh结尾。

	Uri string `json:"uri"`
	// 引导操作脚本参数。

	Parameters *string `json:"parameters,omitempty"`
	// 引导操作脚本所执行的节点类型，包含Master、Core和Task三种类型。

	Nodes []string `json:"nodes"`
	// 引导操作脚本是否只运行在主Master节点上。 缺省值为false，表示引导操作脚本可运行在所有Master节点上。

	ActiveMaster *bool `json:"active_master,omitempty"`
	// 引导操作脚本执行失败后，是否继续执行后续脚本和创建集群。  缺省值为errorout，表示终止操作。   说明： 建议您在调试阶段设置为“继续”，无论此引导操作是否执行成功，则集群都能继续安装和启动。  枚举值： - continue：继续执行后续脚本。 - errorout：终止操作。

	FailAction BootstrapScriptRespFailAction `json:"fail_action"`
	// 引导操作脚本执行的时间。目前支持“组件启动前”和“组件启动后”两种类型。 缺省值为false，表示引导操作脚本在组件启动后执行。

	BeforeComponentStart *bool `json:"before_component_start,omitempty"`
	// 单个引导操作脚本的执行时间。

	StartTime *int64 `json:"start_time,omitempty"`
	// 单个引导操作脚本的运行状态。 - PENDING - IN_PROGRESS - SUCCESS - FAILURE

	State *string `json:"state,omitempty"`
}

func (o BootstrapScriptResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BootstrapScriptResp struct{}"
	}

	return strings.Join([]string{"BootstrapScriptResp", string(data)}, " ")
}

type BootstrapScriptRespFailAction struct {
	value string
}

type BootstrapScriptRespFailActionEnum struct {
	CONTINUE BootstrapScriptRespFailAction
	ERROROUT BootstrapScriptRespFailAction
}

func GetBootstrapScriptRespFailActionEnum() BootstrapScriptRespFailActionEnum {
	return BootstrapScriptRespFailActionEnum{
		CONTINUE: BootstrapScriptRespFailAction{
			value: "continue",
		},
		ERROROUT: BootstrapScriptRespFailAction{
			value: "errorout",
		},
	}
}

func (c BootstrapScriptRespFailAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BootstrapScriptRespFailAction) UnmarshalJSON(b []byte) error {
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
