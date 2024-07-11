package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteHostsRequest Request Object
type BatchDeleteHostsRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	Body *DeploymentHostListEntity `json:"body,omitempty"`
}

func (o BatchDeleteHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteHostsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteHostsRequest", string(data)}, " ")
}
