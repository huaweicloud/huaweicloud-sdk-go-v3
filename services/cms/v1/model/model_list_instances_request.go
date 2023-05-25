package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListInstancesRequest struct {

	// 查询返回实例数量限制 取值范围：1-1000。
	Limit *int32 `json:"limit,omitempty"`

	// 取值为上一页数据的最后一条记录的唯一标识
	Marker *string `json:"marker,omitempty"`

	// 智能购买组id
	AutoLaunchGroupId string `json:"auto_launch_group_id"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}
