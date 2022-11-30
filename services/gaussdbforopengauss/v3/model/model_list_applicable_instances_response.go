package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApplicableInstancesResponse struct {

	// 实例列表，显示实例ID和实例名称。
	Instances *[]InstancesListResult `json:"instances,omitempty"`

	// 查询数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListApplicableInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicableInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListApplicableInstancesResponse", string(data)}, " ")
}
