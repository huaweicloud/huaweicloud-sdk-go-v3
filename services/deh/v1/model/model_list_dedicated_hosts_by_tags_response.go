package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDedicatedHostsByTagsResponse struct {
	// 返回的专属主机列表。

	Resources *[]RespDeh `json:"resources,omitempty"`
	// 总记录数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDedicatedHostsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostsByTagsResponse", string(data)}, " ")
}
