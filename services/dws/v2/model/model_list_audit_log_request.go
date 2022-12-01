package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAuditLogRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListAuditLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditLogRequest struct{}"
	}

	return strings.Join([]string{"ListAuditLogRequest", string(data)}, " ")
}
