package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeploymentResources 创建容器时使用的资源
type DeploymentResources struct {

	// 允许容器使用的最大资源，key值支持填写：cpu, memory, gpu, npu, D310, D910
	Limits map[string]string `json:"limits,omitempty"`

	// 容器需要使用的最小资源，key值支持填写：cpu, memory, gpu, npu, D310, D910
	Requests map[string]string `json:"requests,omitempty"`
}

func (o DeploymentResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentResources struct{}"
	}

	return strings.Join([]string{"DeploymentResources", string(data)}, " ")
}
