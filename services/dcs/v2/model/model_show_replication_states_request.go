package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplicationStatesRequest Request Object
type ShowReplicationStatesRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 分片ID。
	GroupId string `json:"group_id"`
}

func (o ShowReplicationStatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationStatesRequest struct{}"
	}

	return strings.Join([]string{"ShowReplicationStatesRequest", string(data)}, " ")
}
