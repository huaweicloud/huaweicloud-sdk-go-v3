package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JoinGroupRequest Request Object
type JoinGroupRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`

	Body *ClusterJoinGroupRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o JoinGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JoinGroupRequest struct{}"
	}

	return strings.Join([]string{"JoinGroupRequest", string(data)}, " ")
}
