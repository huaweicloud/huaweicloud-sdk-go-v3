package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListResourceByTagsResponse struct {
	// 资源实例

	Resources *[]Resource `json:"resources,omitempty"`
	// 总记录数

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourceByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceByTagsResponse", string(data)}, " ")
}
