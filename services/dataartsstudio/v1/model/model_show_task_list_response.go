package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskListResponse Response Object
type ShowTaskListResponse struct {

	// 查询采集任务数量
	Count *int32 `json:"count,omitempty"`

	// 同一projectId下已创建采集任务数量
	TotalRecords *int32 `json:"total_records,omitempty"`

	// 同一projectId下允许创建采集任务数量
	MaxRecords *int32 `json:"max_records,omitempty"`

	// 采集任务列表
	Resources      *[]MetadataCollectionTask `json:"resources,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowTaskListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskListResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskListResponse", string(data)}, " ")
}
