package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModelGroupsResponse Response Object
type ListModelGroupsResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 模型组列表项。
	Items          *[]ModelGroupItemVo `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListModelGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModelGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListModelGroupsResponse", string(data)}, " ")
}
