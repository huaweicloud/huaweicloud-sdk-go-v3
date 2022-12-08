package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAuditInstanceJobsResponse struct {

	// 实例创建任务列表
	Jobs           *[]JobBean `json:"jobs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListAuditInstanceJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditInstanceJobsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditInstanceJobsResponse", string(data)}, " ")
}
