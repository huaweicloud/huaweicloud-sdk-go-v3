package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubTaskCaseOverstockUsingRequest Request Object
type ListSubTaskCaseOverstockUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	// 数据开始时间
	StartTime int64 `json:"startTime"`

	// 数据结束时间
	EndTime int64 `json:"endTime"`

	// 执行机类型
	ExecutorType *string `json:"executorType,omitempty"`

	// 执行机标签
	Label string `json:"label"`

	// 执行机所属区域id
	LocationId *string `json:"locationId,omitempty"`

	// 分页当前页码
	PageNum *int32 `json:"pageNum,omitempty"`

	// 分页大小(分页参数只要有一个为空即不做分页)
	PageSize *int32 `json:"pageSize,omitempty"`
}

func (o ListSubTaskCaseOverstockUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubTaskCaseOverstockUsingRequest struct{}"
	}

	return strings.Join([]string{"ListSubTaskCaseOverstockUsingRequest", string(data)}, " ")
}
