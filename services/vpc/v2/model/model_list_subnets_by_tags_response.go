package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSubnetsByTagsResponse struct {

	// 资源列表
	Resources *[]ListResourceResp `json:"resources,omitempty" xml:"resources"`

	// 资源数量
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSubnetsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSubnetsByTagsResponse", string(data)}, " ")
}
