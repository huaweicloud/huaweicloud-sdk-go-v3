package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestorableInstancesResponse Response Object
type ListRestorableInstancesResponse struct {

	// 返回可用于备份恢复的实例列表。
	Instances *[]InstancesResult `json:"instances,omitempty"`

	// 查询出来的实例总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRestorableInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestorableInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListRestorableInstancesResponse", string(data)}, " ")
}
