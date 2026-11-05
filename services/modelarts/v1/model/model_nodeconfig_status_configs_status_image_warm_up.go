package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NodeconfigStatusConfigsStatusImageWarmUp **参数解释**：镜像预热任务的执行状态。 **约束限制**：仅在请求体/插件中存在镜像预热配置（image_warm_up）时返回。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeconfigStatusConfigsStatusImageWarmUp struct {

	// **参数解释**：镜像预热任务的整体状态。 **约束限制**：不涉及。 **取值范围**：可选值如下： - Success：全部成功，所有镜像在所有节点上均预热成功。 - Failed：全部失败，所有镜像在所有节点上均预热失败。 - Pulling：预热中，镜像正在被拉取。 - PartialCompletion：部分完成，部分镜像成功、部分失败或部分节点成功。 - Unknown：未知状态，解析节点上报数据异常。 **默认取值**：不涉及。
	State *NodeconfigStatusConfigsStatusImageWarmUpState `json:"state,omitempty"`

	// **参数解释**：镜像预热任务的详细结果信息（JSON格式字符串），包含以下字段： - successfulImages：预热成功的镜像列表 - failedImages：预热失败的镜像列表 - partialSuccess：部分节点成功的镜像列表 - pendingImages：尚未完成预热的镜像列表 - additionalInfos：失败原因等附加信息（如镜像不存在、拉取超时、拉取失败、镜像仓库中manifest不存在） - successfulImagesConfig：成功镜像的配置详情（id、name、creationTimestamp、type、image） - failedImagesConfig：失败镜像的配置详情 - partialSuccessConfig：部分成功镜像的配置详情 - pendingImagesConfig：待预热镜像的配置详情 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Message *string `json:"message,omitempty"`
}

func (o NodeconfigStatusConfigsStatusImageWarmUp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigStatusConfigsStatusImageWarmUp struct{}"
	}

	return strings.Join([]string{"NodeconfigStatusConfigsStatusImageWarmUp", string(data)}, " ")
}

type NodeconfigStatusConfigsStatusImageWarmUpState struct {
	value string
}

type NodeconfigStatusConfigsStatusImageWarmUpStateEnum struct {
	SUCCESS            NodeconfigStatusConfigsStatusImageWarmUpState
	FAILED             NodeconfigStatusConfigsStatusImageWarmUpState
	PULLING            NodeconfigStatusConfigsStatusImageWarmUpState
	PARTIAL_COMPLETION NodeconfigStatusConfigsStatusImageWarmUpState
	UNKNOWN            NodeconfigStatusConfigsStatusImageWarmUpState
}

func GetNodeconfigStatusConfigsStatusImageWarmUpStateEnum() NodeconfigStatusConfigsStatusImageWarmUpStateEnum {
	return NodeconfigStatusConfigsStatusImageWarmUpStateEnum{
		SUCCESS: NodeconfigStatusConfigsStatusImageWarmUpState{
			value: "Success",
		},
		FAILED: NodeconfigStatusConfigsStatusImageWarmUpState{
			value: "Failed",
		},
		PULLING: NodeconfigStatusConfigsStatusImageWarmUpState{
			value: "Pulling",
		},
		PARTIAL_COMPLETION: NodeconfigStatusConfigsStatusImageWarmUpState{
			value: "PartialCompletion",
		},
		UNKNOWN: NodeconfigStatusConfigsStatusImageWarmUpState{
			value: "Unknown",
		},
	}
}

func (c NodeconfigStatusConfigsStatusImageWarmUpState) Value() string {
	return c.value
}

func (c NodeconfigStatusConfigsStatusImageWarmUpState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeconfigStatusConfigsStatusImageWarmUpState) UnmarshalJSON(b []byte) error {
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
