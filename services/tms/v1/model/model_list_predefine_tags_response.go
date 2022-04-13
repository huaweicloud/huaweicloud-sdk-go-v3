package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPredefineTagsResponse struct {
	// 分页位置标识（索引）。

	Marker *string `json:"marker,omitempty"`
	// 查询到的标签总数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 查询到的标签列表

	Tags           *[]PredefineTag `json:"tags,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListPredefineTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPredefineTagsResponse struct{}"
	}

	return strings.Join([]string{"ListPredefineTagsResponse", string(data)}, " ")
}
