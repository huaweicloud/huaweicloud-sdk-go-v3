package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomerConfig 用户配置信息
type CustomerConfig struct {

	// 失败提醒。
	FailureRemind *string `json:"failureRemind,omitempty"`

	// 集群类型。
	ClusterName *string `json:"clusterName,omitempty"`

	// 服务提供
	ServiceProvider *string `json:"serviceProvider,omitempty"`

	// 是否本地磁盘。
	LocalDisk *string `json:"localDisk,omitempty"`

	// 是否使用ssl。
	Ssl *string `json:"ssl,omitempty"`

	// 创建来源
	CreateFrom *string `json:"createFrom,omitempty"`

	// 资源ID
	ResourceId *string `json:"resourceId,omitempty"`

	// 规格类型
	FlavorType *string `json:"flavorType,omitempty"`

	// 工作空间ID
	WorkSpaceId *string `json:"workSpaceId,omitempty"`

	// 适用
	Trial *string `json:"trial,omitempty"`
}

func (o CustomerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerConfig struct{}"
	}

	return strings.Join([]string{"CustomerConfig", string(data)}, " ")
}
