package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddFullSqlTasksRequest Request Object
type BatchAddFullSqlTasksRequest struct {
	Body *BatchAddFullSqlTasksRequestBody `json:"body,omitempty"`
}

func (o BatchAddFullSqlTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddFullSqlTasksRequest struct{}"
	}

	return strings.Join([]string{"BatchAddFullSqlTasksRequest", string(data)}, " ")
}
