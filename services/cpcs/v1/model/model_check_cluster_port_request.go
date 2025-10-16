package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckClusterPortRequest Request Object
type CheckClusterPortRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 端口的id
	Id string `json:"id"`
}

func (o CheckClusterPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckClusterPortRequest struct{}"
	}

	return strings.Join([]string{"CheckClusterPortRequest", string(data)}, " ")
}
