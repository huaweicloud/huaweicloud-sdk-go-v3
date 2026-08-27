package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplyObjectsResponse Response Object
type ListApplyObjectsResponse struct {

	// 总数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 应用对象列表
	ApplyObjects   *[]ApplyObjectDetailInfo `json:"apply_objects,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListApplyObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplyObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListApplyObjectsResponse", string(data)}, " ")
}
