package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountInstanceByTagsResponse Response Object
type CountInstanceByTagsResponse struct {

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountInstanceByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountInstanceByTagsResponse struct{}"
	}

	return strings.Join([]string{"CountInstanceByTagsResponse", string(data)}, " ")
}
