package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostClusterDetailRequest Request Object
type ShowHostClusterDetailRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`
}

func (o ShowHostClusterDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostClusterDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowHostClusterDetailRequest", string(data)}, " ")
}
