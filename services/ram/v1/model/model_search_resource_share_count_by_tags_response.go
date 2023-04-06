package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchResourceShareCountByTagsResponse struct {

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SearchResourceShareCountByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchResourceShareCountByTagsResponse struct{}"
	}

	return strings.Join([]string{"SearchResourceShareCountByTagsResponse", string(data)}, " ")
}
