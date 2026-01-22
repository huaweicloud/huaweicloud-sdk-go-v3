package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListReferencesResponse Response Object
type ListReferencesResponse struct {

	// 下一次分页查询的起始ID。如果未返回该值，则表示数据已查询完毕
	NextMarker *string `json:"next_marker,omitempty"`

	// 分页查询时表示是否还有下一页的数据
	HasMore *bool `json:"has_more,omitempty"`

	// 镜像版本列表
	Tags           *[]string `json:"tags,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListReferencesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReferencesResponse struct{}"
	}

	return strings.Join([]string{"ListReferencesResponse", string(data)}, " ")
}
