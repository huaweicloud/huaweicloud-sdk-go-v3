package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostClusterRequest Request Object
type UpdateHostClusterRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	Body *HostClusterRequest `json:"body,omitempty"`
}

func (o UpdateHostClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostClusterRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostClusterRequest", string(data)}, " ")
}
