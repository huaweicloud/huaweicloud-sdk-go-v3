package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityGroupsByTagsResponse Response Object
type ListSecurityGroupsByTagsResponse struct {

	// 资源列表
	Resources *[]ListResourceResp `json:"resources,omitempty"`

	// 资源数量
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSecurityGroupsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupsByTagsResponse", string(data)}, " ")
}
