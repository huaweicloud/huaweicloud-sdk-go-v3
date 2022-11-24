package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentSpec struct {

	// 语言/运行时，例如：Java8、tomcat8。
	Runtime *string `json:"runtime,omitempty"`

	// 环境id。
	EnvId *string `json:"env_id,omitempty"`

	// 副本。
	Replica *int32 `json:"replica,omitempty"`

	// 可用副本。
	AvailableReplica *int32 `json:"available_replica,omitempty"`

	Source *Source `json:"source,omitempty"`

	Build *Build `json:"build,omitempty"`

	AccessInfo *[]Access `json:"access_info,omitempty"`

	// 构建id。
	BuildId *string `json:"build_id,omitempty"`

	// 镜像地址。
	ImageUrl *string `json:"image_url,omitempty"`

	// 任务id。
	JobId *string `json:"job_id,omitempty"`

	// 日志策略。
	LogStrategy *[]LogStrategy `json:"log_strategy,omitempty"`

	// 组件状态
	Status *string `json:"status,omitempty"`
}

func (o ComponentSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentSpec struct{}"
	}

	return strings.Join([]string{"ComponentSpec", string(data)}, " ")
}
