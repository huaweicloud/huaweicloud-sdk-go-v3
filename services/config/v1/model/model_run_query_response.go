package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunQueryResponse Response Object
type RunQueryResponse struct {
	QueryInfo *QueryInfo `json:"query_info,omitempty"`

	// ResourceQL 查询结果.
	Results        *[]interface{} `json:"results,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o RunQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryResponse struct{}"
	}

	return strings.Join([]string{"RunQueryResponse", string(data)}, " ")
}
