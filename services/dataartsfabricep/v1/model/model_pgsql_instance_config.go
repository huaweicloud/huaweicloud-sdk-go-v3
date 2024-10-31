package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PgsqlInstanceConfig PGSQL服务实例的配置，在DeployService中使用
type PgsqlInstanceConfig struct {

	// 资源规格，从规格列表查询获取。
	CoordinatorResourceSpec string `json:"coordinator_resource_spec"`

	// Coordinator的POD标签
	CoordinatorPodLabel string `json:"coordinator_pod_label"`

	// 资源规格，从规格列表查询获取。
	WorkerResourceSpec string `json:"worker_resource_spec"`

	// Worker的POD标签
	WorkerPodLabel string `json:"worker_pod_label"`

	// Coordinator对外服务的端口
	CoordinatorPort int32 `json:"coordinator_port"`

	// json格式 元数据以及data位置信息
	RuntimeParam string `json:"runtime_param"`
}

func (o PgsqlInstanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PgsqlInstanceConfig struct{}"
	}

	return strings.Join([]string{"PgsqlInstanceConfig", string(data)}, " ")
}
