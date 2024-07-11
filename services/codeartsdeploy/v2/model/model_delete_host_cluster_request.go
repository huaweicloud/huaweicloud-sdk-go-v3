package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostClusterRequest Request Object
type DeleteHostClusterRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`
}

func (o DeleteHostClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteHostClusterRequest", string(data)}, " ")
}
