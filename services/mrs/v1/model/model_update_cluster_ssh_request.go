package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterSshRequest Request Object
type UpdateClusterSshRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 开启/关闭集群节点授权
	Enable bool `json:"enable"`

	// 开启集群节点授权截止时间，仅enable为true时需要传递，如不传该值默认为1天授权时间
	ExpireTime float32 `json:"expire_time,omitempty"`
}

func (o UpdateClusterSshRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterSshRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterSshRequest", string(data)}, " ")
}
