package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFsTasksRequest Request Object
type ListFsTasksRequest struct {

	// MIME类型, application/json
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 任务类型。例，DU任务取值为dir-usage
	Feature string `json:"feature"`

	// marker填写为上一页返回的最后一条记录的task_id，作为下一次分页的起始值。比如，marker为d651ea2b-2b20-4c6d-8bbf-2adcec18dac9，limit为100，表示从d651ea2b-2b20-4c6d-8bbf-2adcec18dac9后返回最多100条满足查询条件的记录。
	Marker *string `json:"marker,omitempty"`

	// limit, 取值为正整数. 默认为20，最大值为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFsTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFsTasksRequest struct{}"
	}

	return strings.Join([]string{"ListFsTasksRequest", string(data)}, " ")
}
