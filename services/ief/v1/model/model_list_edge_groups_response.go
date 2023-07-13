package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeGroupsResponse Response Object
type ListEdgeGroupsResponse struct {

	// 边缘节点组数目
	Count *int32 `json:"count,omitempty"`

	// 边缘节点组详情
	EdgeGroups     *[]EdgeGroupResp `json:"edge_groups,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListEdgeGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeGroupsResponse", string(data)}, " ")
}
