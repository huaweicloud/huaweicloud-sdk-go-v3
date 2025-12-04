package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWarmPoolInstancesNewRequest Request Object
type ListWarmPoolInstancesNewRequest struct {

	// 伸缩组ID
	ScalingGroupId string `json:"scaling_group_id"`

	// 查询的记录条数，不传默认20，最大可传入100
	Limit *int32 `json:"limit,omitempty"`

	// 查询暖池实例的分页marker
	Marker *string `json:"marker,omitempty"`
}

func (o ListWarmPoolInstancesNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWarmPoolInstancesNewRequest struct{}"
	}

	return strings.Join([]string{"ListWarmPoolInstancesNewRequest", string(data)}, " ")
}
