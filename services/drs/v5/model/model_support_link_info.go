package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupportLinkInfo 支持的链路信息
type SupportLinkInfo struct {

	// 引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 网络类型。取值： - eip：公网网络。 - vpc：VPC网络，灾备场景不支持选择VPC网络。 - vpn：VPN、专线网络。
	NetType *string `json:"net_type,omitempty"`

	// 迁移模式。取值： - FULL_TRANS ：全量。 - FULL_INCR_TRANS：全量+增量。 - INCR_TRANS：增量。
	TaskModes *[]string `json:"task_modes,omitempty"`

	// 迁移方向。取值： - up：入云 ，灾备场景时对应本云为备。 - down：出云，灾备场景时对应本云为主。 - non-dbs：自建。
	JobDirection *string `json:"job_direction,omitempty"`

	// 云上实例类型。取值： - Single：单机模式。 - Ha：主备模式。 - Cluster：集群模式。 - Sharding：分片模式。 - Independent：GaussDB独立部署模式。
	ClusterMode *string `json:"cluster_mode,omitempty"`

	// DRS实例类型。取值： - Single ：单机。 - Ha：主备。
	JobInstanceType *string `json:"job_instance_type,omitempty"`

	// 是否支持绑定EIP
	IsSupportBindEip *bool `json:"is_support_bind_eip,omitempty"`
}

func (o SupportLinkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupportLinkInfo struct{}"
	}

	return strings.Join([]string{"SupportLinkInfo", string(data)}, " ")
}
