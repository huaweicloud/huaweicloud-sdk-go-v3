package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户配置信息
type CustomerConfig struct {

	// 失败提醒。
	FailureRemind *string `json:"failureRemind,omitempty" xml:"failureRemind"`

	// 集群类型。
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName"`

	// 服务提供
	ServiceProvider *string `json:"serviceProvider,omitempty" xml:"serviceProvider"`

	// 是否本地磁盘。
	LocalDisk *string `json:"localDisk,omitempty" xml:"localDisk"`

	// 是否使用ssl。
	Ssl *string `json:"ssl,omitempty" xml:"ssl"`

	// 创建来源
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom"`

	// 资源ID
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId"`

	// 规格类型
	FlavorType *string `json:"flavorType,omitempty" xml:"flavorType"`

	// 工作空间ID
	WorkSpaceId *string `json:"workSpaceId,omitempty" xml:"workSpaceId"`

	// 适用
	Trial *string `json:"trial,omitempty" xml:"trial"`
}

func (o CustomerConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerConfig struct{}"
	}

	return strings.Join([]string{"CustomerConfig", string(data)}, " ")
}
