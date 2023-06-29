package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditInstanceJobsRequest Request Object
type ListAuditInstanceJobsRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`
}

func (o ListAuditInstanceJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditInstanceJobsRequest struct{}"
	}

	return strings.Join([]string{"ListAuditInstanceJobsRequest", string(data)}, " ")
}
