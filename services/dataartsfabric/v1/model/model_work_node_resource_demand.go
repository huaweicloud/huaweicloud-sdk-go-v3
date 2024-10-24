package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkNodeResourceDemand ray集群工作节点资源需求量
type WorkNodeResourceDemand struct {

	// 资源配置名称
	Name string `json:"name"`

	Resource *ResourceDemand `json:"resource"`
}

func (o WorkNodeResourceDemand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkNodeResourceDemand struct{}"
	}

	return strings.Join([]string{"WorkNodeResourceDemand", string(data)}, " ")
}
