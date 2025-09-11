package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditInstanceJobsNewRequest Request Object
type ListAuditInstanceJobsNewRequest struct {

	// 资源ID。可在查询实例列表接口的resource_id获取。
	ResourceId string `json:"resource_id"`
}

func (o ListAuditInstanceJobsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditInstanceJobsNewRequest struct{}"
	}

	return strings.Join([]string{"ListAuditInstanceJobsNewRequest", string(data)}, " ")
}
