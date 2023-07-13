package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisasterRecoveryCluster 容灾集群信息
type DisasterRecoveryCluster struct {

	// 容灾集群信息ID
	Id *string `json:"id,omitempty"`

	// 容灾集群名称
	Name *string `json:"name,omitempty"`

	// 容灾集群所在AZ
	ClusterAz *string `json:"cluster_az,omitempty"`

	// 容灾集群角色
	Role *string `json:"role,omitempty"`

	// 容灾集群所在region
	Region *string `json:"region,omitempty"`

	// 容灾集群状态
	Status *string `json:"status,omitempty"`

	// 容灾进度
	Progress *string `json:"progress,omitempty"`

	// 上一次容灾时间
	LastSuccessTime *string `json:"last_success_time,omitempty"`

	// OBS桶名称
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`
}

func (o DisasterRecoveryCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecoveryCluster struct{}"
	}

	return strings.Join([]string{"DisasterRecoveryCluster", string(data)}, " ")
}
