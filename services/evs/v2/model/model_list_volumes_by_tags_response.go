package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumesByTagsResponse Response Object
type ListVolumesByTagsResponse struct {

	// 符合查询条件的云硬盘资源个数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 符合查询条件的资源列表
	Resources      *[]Resource `json:"resources,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListVolumesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListVolumesByTagsResponse", string(data)}, " ")
}
