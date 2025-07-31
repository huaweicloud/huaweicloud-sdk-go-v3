package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerlessJobInfo 普通任务基本信息
type ServerlessJobInfo struct {

	// 普通任务名称
	Name *string `json:"name,omitempty"`

	// 命名空间名称
	NamespaceName *string `json:"namespace_name,omitempty"`

	// 所属集群
	ClusterName *string `json:"cluster_name,omitempty"`

	// 状态，包含以下几种 -Running：正常运行 -Failed：存在异常
	Status *string `json:"status,omitempty"`

	// 实例个数
	PodsNum *int32 `json:"pods_num,omitempty"`

	// 镜像名称
	ImageName *string `json:"image_name,omitempty"`

	// 标签
	MatchLabels *[]LabelInfo `json:"match_labels,omitempty"`

	// 执行时间
	ExecuteTime *int64 `json:"execute_time,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o ServerlessJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerlessJobInfo struct{}"
	}

	return strings.Join([]string{"ServerlessJobInfo", string(data)}, " ")
}
