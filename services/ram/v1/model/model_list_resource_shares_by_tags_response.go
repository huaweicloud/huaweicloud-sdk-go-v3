package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResourceSharesByTagsResponse struct {

	// 资源信息列表。
	Resources *[]ResourceDto `json:"resources,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourceSharesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceSharesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceSharesByTagsResponse", string(data)}, " ")
}
