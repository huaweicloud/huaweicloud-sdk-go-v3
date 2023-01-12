package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListGraphs2Response struct {

	// 图总个数。请求失败时为空。
	GraphCount *int32 `json:"graph_count,omitempty"`

	// 图列表。请求失败时为空。
	Graphs         *[]ListGraphsRespGraphs `json:"graphs,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListGraphs2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphs2Response struct{}"
	}

	return strings.Join([]string{"ListGraphs2Response", string(data)}, " ")
}
