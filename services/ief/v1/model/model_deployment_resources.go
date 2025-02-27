package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeploymentResources 创建容器时使用的资源
type DeploymentResources struct {

	// - 允许容器使用的最大资源，key值支持填写：cpu, memory, gpu, npu。 - 申请NPU资源时可填写指定的NPU芯片类型，支持填写：D310、D310B。注意：key值填写为npu时，默认使用npu_type字段指定的芯片类型。如果npu_type字段为空，默认使用D310芯片类型。
	Limits map[string]string `json:"limits,omitempty"`

	// - 容器需要使用的最小资源，key值支持填写：cpu, memory, gpu, npu - 申请NPU资源时可填写指定的NPU芯片类型，支持填写：D310、D310B。注意：key值填写为npu时，默认使用npu_type字段指定的芯片类型。如果npu_type字段为空，默认使用D310芯片类型。
	Requests map[string]string `json:"requests,omitempty"`
}

func (o DeploymentResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentResources struct{}"
	}

	return strings.Join([]string{"DeploymentResources", string(data)}, " ")
}
