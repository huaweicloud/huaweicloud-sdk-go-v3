package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceGroupsResponse Response Object
type BatchDeleteResourceGroupsResponse struct {

	// 成功删除的资源分组ID列表
	GroupIds       *[]string `json:"group_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteResourceGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceGroupsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceGroupsResponse", string(data)}, " ")
}
