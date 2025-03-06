package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDwsClusterRequest Request Object
type DeleteDwsClusterRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 集群需要保留的快照数
	KeepLastManualBackup *string `json:"keep_last_manual_backup,omitempty"`

	// 集群是否释放弹性公网IP，默认是NO_RELEASE不释放绑定的弹性公网IP，RELEASE_BINDING释放绑定的弹性公网IP
	ReleaseEipType *string `json:"release_eip_type,omitempty"`
}

func (o DeleteDwsClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDwsClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteDwsClusterRequest", string(data)}, " ")
}
