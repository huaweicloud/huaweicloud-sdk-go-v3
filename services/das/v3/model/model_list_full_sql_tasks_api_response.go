package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFullSqlTasksApiResponse Response Object
type ListFullSqlTasksApiResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// SQL洞察任务明细
	Tasks          *[]SqlParseTask `json:"tasks,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListFullSqlTasksApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullSqlTasksApiResponse struct{}"
	}

	return strings.Join([]string{"ListFullSqlTasksApiResponse", string(data)}, " ")
}
