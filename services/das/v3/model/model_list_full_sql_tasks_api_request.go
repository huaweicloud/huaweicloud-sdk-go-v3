package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFullSqlTasksApiRequest Request Object
type ListFullSqlTasksApiRequest struct {
	Body *ListFullSqlTasksRequestBody `json:"body,omitempty"`
}

func (o ListFullSqlTasksApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullSqlTasksApiRequest struct{}"
	}

	return strings.Join([]string{"ListFullSqlTasksApiRequest", string(data)}, " ")
}
