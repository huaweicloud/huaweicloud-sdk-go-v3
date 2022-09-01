package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPublicipsByTagsResponse struct {

	// resource对象列表
	Resources *[]ListResourceResp `json:"resources,omitempty" xml:"resources"`

	// 总记录数
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPublicipsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListPublicipsByTagsResponse", string(data)}, " ")
}
