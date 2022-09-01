package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDedicatedHostsByTagsResponse struct {

	// 返回的专属主机列表。
	Resources *[]RespDeh `json:"resources,omitempty" xml:"resources"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDedicatedHostsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostsByTagsResponse", string(data)}, " ")
}
