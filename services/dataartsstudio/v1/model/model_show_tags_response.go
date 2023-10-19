package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTagsResponse Response Object
type ShowTagsResponse struct {

	// 每页大小
	Limit *int32 `json:"limit,omitempty"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 查询出来的条数
	Count *int32 `json:"count,omitempty"`

	// 可创建标签数量配额额
	Quota *int32 `json:"quota,omitempty"`

	// 标签实体
	Tags *[]OpenTag `json:"tags,omitempty"`

	// 已创建的标签总条数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowTagsResponse", string(data)}, " ")
}
