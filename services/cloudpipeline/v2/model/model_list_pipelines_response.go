package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPipelinesResponse struct {

	// 起始偏移
	Offset *int32 `json:"offset,omitempty"`

	// 查询大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 流水线
	Pipelines      *[]ListPipelinesPagePipelines `json:"pipelines,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListPipelinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesResponse struct{}"
	}

	return strings.Join([]string{"ListPipelinesResponse", string(data)}, " ")
}
