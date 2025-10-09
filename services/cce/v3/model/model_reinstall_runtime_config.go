package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReinstallRuntimeConfig 节点重装场景容器运行时配置
type ReinstallRuntimeConfig struct {

	// **参数解释**: 节点上单容器的可用磁盘空间大小（已废弃，请优先使用containerBaseSize参数），单位G。  不配置该值或值为0时将使用默认值，Devicemapper模式下默认值为10；OverlayFS模式默认不限制单容器可用空间大小，且dockerBaseSize设置仅在新版本集群的EulerOS[/HCEOS2.0](tag:hws,hws_hk,ctc,cmcc)节点上生效。  CCE节点容器运行时空间配置请参考[数据盘空间分配说明](cce_01_0341.xml)。  Devicemapper模式下建议dockerBaseSize配置不超过80G，设置过大时可能会导致容器运行时初始化时间过长而启动失败，若对容器磁盘大小有特殊要求，可考虑使用挂载外部或本地存储方式代替。 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
	DockerBaseSize *int32 `json:"dockerBaseSize,omitempty"`

	// **参数解释**: 节点上单容器的可用磁盘空间大小，单位G。  不配置该值或值为0时将使用默认值，OverlayFS模式默认不限制单容器可用空间大小；Devicemapper模式下默认值为10，且containerBaseSize设置仅在新版本集群（v1.23.14-r0/v1.25.9-r0/v1.27.6-r0/v1.28.4-r0及以上）的EulerOS[/HCEOS2.0](tag:hws,hws_hk,ctc,cmcc)节点上生效。  CCE节点容器运行时空间配置请参考[数据盘空间分配说明](cce_01_0341.xml)。  Devicemapper模式下建议containerBaseSize配置不超过80G，设置过大时可能会导致容器运行时初始化时间过长而启动失败，若对容器磁盘大小有特殊要求，可考虑使用挂载外部或本地存储方式代替；Devicemapper模式在新版中仅有共池裸机使用，已逐步废弃。  **约束限制**: > 更新节点池时，不支持更新此参数  **取值范围**: 不涉及 **默认取值**: 不涉及
	ContainerBaseSize *int32 `json:"containerBaseSize,omitempty"`

	Runtime *Runtime `json:"runtime,omitempty"`
}

func (o ReinstallRuntimeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallRuntimeConfig struct{}"
	}

	return strings.Join([]string{"ReinstallRuntimeConfig", string(data)}, " ")
}
