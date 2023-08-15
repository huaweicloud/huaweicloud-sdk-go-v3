package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBatchesRequest Request Object
type ListBatchesRequest struct {

	// DLI队列名称，不填写则获取当前Project下所有批处理作业(不推荐使用)。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 用于查询开始时间在该时间点之前的作业。时间格式为unix时间戳，单位：毫秒。
	End *int64 `json:"end,omitempty"`

	// 起始批处理作业的索引号，默认从0开始。
	From *int32 `json:"from,omitempty"`

	JobId *string `json:"job-id,omitempty"`

	// 指定作业排序方式，默认为CREATE_TIME_DESC（作业提交时间降序），支持DURATION_DESC（作业运行时长降序）、DURATION_ASC（作业运行时长升序）、CREATE_TIME_DESC（作业提交时间降序）、CREATE_TIME_ASC（作业提交时间升序）四种排序方式。
	Order *string `json:"order,omitempty"`

	QueueName *string `json:"queue_name,omitempty"`

	// 查询批处理作业的数量。
	Size *int32 `json:"size,omitempty"`

	// 用于查询开始时间在该时间点之后的作业。时间格式为unix时间戳，单位：毫秒。
	Start *int64 `json:"start,omitempty"`

	State *string `json:"state,omitempty"`
}

func (o ListBatchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBatchesRequest struct{}"
	}

	return strings.Join([]string{"ListBatchesRequest", string(data)}, " ")
}
