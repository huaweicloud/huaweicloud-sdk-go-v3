package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogsJobResponse Response Object
type ListLogsJobResponse struct {
	ClusterLogRecord *[]ClusterLogRecord `json:"clusterLogRecord,omitempty"`

	// **参数解释**： 日志记录总条数。 **取值范围**： 不涉及
	TotalSize      *int32 `json:"totalSize,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLogsJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogsJobResponse struct{}"
	}

	return strings.Join([]string{"ListLogsJobResponse", string(data)}, " ")
}
