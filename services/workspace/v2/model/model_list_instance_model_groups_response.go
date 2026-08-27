package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceModelGroupsResponse Response Object
type ListInstanceModelGroupsResponse struct {

	// 总数。
	Total *int32 `json:"total,omitempty"`

	// 模型分组列表，按优先级升序排列。
	Items          *[]InstanceModelGroupItem `json:"items,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListInstanceModelGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceModelGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceModelGroupsResponse", string(data)}, " ")
}
