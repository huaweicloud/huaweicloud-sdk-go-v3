package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 邀请者的信息
type InvitorInfo struct {
	// 邀请方用户id

	InvitorId *string `json:"invitor_id,omitempty"`
	// 邀请方用户名

	InvitorName *string `json:"invitor_name,omitempty"`
	// 邀请方的服务id

	InvitorBcsId *string `json:"invitor_bcs_id,omitempty"`
	// 邀请方的服务名

	InvitorBcsName *string `json:"invitor_bcs_name,omitempty"`
	// 邀请方的共识

	InvitorConsensus *string `json:"invitor_consensus,omitempty"`
	// 邀请方的projectID

	InvitorProjectId *string `json:"invitor_project_id,omitempty"`
	// 邀请方的集群类型

	InvitorClusterType *string `json:"invitor_cluster_type,omitempty"`
	// 邀请方数据库类型

	InvitorDatabaseType *string `json:"invitor_database_type,omitempty"`
	// 邀请方的签名算法

	InvitorSignatureAlgorithm *string `json:"invitor_signature_algorithm,omitempty"`
	// 邀请方的fabric版本

	InvitorFabricVersion *string `json:"invitor_fabric_version,omitempty"`
	// 是否允许order老化

	OrderFadeEnabled *bool `json:"order_fade_enabled,omitempty"`
	// order老化阈值

	OrderFadeCache *int64 `json:"order_fade_cache,omitempty"`
}

func (o InvitorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvitorInfo struct{}"
	}

	return strings.Join([]string{"InvitorInfo", string(data)}, " ")
}
