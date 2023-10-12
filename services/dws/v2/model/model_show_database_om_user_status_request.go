package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseOmUserStatusRequest Request Object
type ShowDatabaseOmUserStatusRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ShowDatabaseOmUserStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseOmUserStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseOmUserStatusRequest", string(data)}, " ")
}
