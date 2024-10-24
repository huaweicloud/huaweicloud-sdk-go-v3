package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkerGroupSpec Ray worker配置信息
type WorkerGroupSpec struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 资源规格，从规格列表查询获取。
	SpecCode *string `json:"spec_code,omitempty"`

	// 最小副本数
	MinReplicas *int32 `json:"min_replicas,omitempty"`

	// 最大副本数
	MaxReplicas *int32 `json:"max_replicas,omitempty"`
}

func (o WorkerGroupSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkerGroupSpec struct{}"
	}

	return strings.Join([]string{"WorkerGroupSpec", string(data)}, " ")
}
