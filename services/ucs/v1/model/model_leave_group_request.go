package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LeaveGroupRequest Request Object
type LeaveGroupRequest struct {

	// 集群id
	Clusterid string `json:"clusterid"`
}

func (o LeaveGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LeaveGroupRequest struct{}"
	}

	return strings.Join([]string{"LeaveGroupRequest", string(data)}, " ")
}
