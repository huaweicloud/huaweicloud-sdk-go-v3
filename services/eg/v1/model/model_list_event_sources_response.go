package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEventSourcesResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items          *[]SourceInfo `json:"items,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListEventSourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventSourcesResponse struct{}"
	}

	return strings.Join([]string{"ListEventSourcesResponse", string(data)}, " ")
}
