package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunAggregateResourceQueryResponse struct {
	QueryInfo *QueryInfo `json:"query_info,omitempty"`

	// ResourceQL 查询结果.
	Results        *[]interface{} `json:"results,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o RunAggregateResourceQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAggregateResourceQueryResponse struct{}"
	}

	return strings.Join([]string{"RunAggregateResourceQueryResponse", string(data)}, " ")
}
