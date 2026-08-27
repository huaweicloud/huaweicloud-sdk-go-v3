package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskApplyObjectsResponse Response Object
type ListTaskApplyObjectsResponse struct {

	// 总数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 应用对象列表
	ApplyObjects   *[]TaskApplyObjectDetailInfo `json:"apply_objects,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListTaskApplyObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskApplyObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListTaskApplyObjectsResponse", string(data)}, " ")
}
