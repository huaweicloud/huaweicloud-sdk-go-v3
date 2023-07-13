package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupReplicationInfoRequest Request Object
type ListGroupReplicationInfoRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ListGroupReplicationInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupReplicationInfoRequest struct{}"
	}

	return strings.Join([]string{"ListGroupReplicationInfoRequest", string(data)}, " ")
}
