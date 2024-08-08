package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesResponse Response Object
type ListInstancesResponse struct {

	// 设备接入实例的总数
	Count *int32 `json:"count,omitempty"`

	// 本次分页查询结果中最后一条记录的ID，可在下一次分页查询时使用
	Marker *string `json:"marker,omitempty"`

	// 设备接入实例的详情列表
	Instances      *[]QueryInstanceSimplify `json:"instances,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesResponse", string(data)}, " ")
}
