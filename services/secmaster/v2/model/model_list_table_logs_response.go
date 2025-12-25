package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableLogsResponse Response Object
type ListTableLogsResponse struct {

	// 查询结果 schema
	Schema *[]SearchQueryField `json:"schema,omitempty"`

	// 查询结果行
	Datarows *[][]interface{} `json:"datarows,omitempty"`

	// Total count of results
	Total *int32 `json:"total,omitempty"`

	// Returned count of results
	Size *int32 `json:"size,omitempty"`

	// Results in JSON format
	Results        *[]SearchQueryResult `json:"results,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListTableLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableLogsResponse struct{}"
	}

	return strings.Join([]string{"ListTableLogsResponse", string(data)}, " ")
}
