package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditInstanceJobsNewResponse Response Object
type ListAuditInstanceJobsNewResponse struct {

	// 实例创建任务列表
	Jobs           *[]JobBean `json:"jobs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListAuditInstanceJobsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditInstanceJobsNewResponse struct{}"
	}

	return strings.Join([]string{"ListAuditInstanceJobsNewResponse", string(data)}, " ")
}
