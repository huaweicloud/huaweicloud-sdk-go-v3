package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExecutionsRequest Request Object
type ListExecutionsRequest struct {

	// 分页参数：每页返回记录个数限制 不传默认第一页查10个
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数：从offset指定的下一条数据开始查询，不传默认0
	Offset *int64 `json:"offset,omitempty"`

	// 创建人，模糊查询
	Creator *string `json:"creator,omitempty"`

	// 开始时间，大于
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间，小于
	EndTime *int64 `json:"end_time,omitempty"`

	// 作业名称，模糊查询
	DocumentName *string `json:"document_name,omitempty"`

	// 作业id
	DocumentId *string `json:"document_id,omitempty"`

	// 标签过滤条件，例：?tags=key1=value1,key2=value2
	Tags *string `json:"tags,omitempty"`

	// 列表查询不返回子工单
	ExcludeChildExecutions *bool `json:"exclude_child_executions,omitempty"`
}

func (o ListExecutionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionsRequest struct{}"
	}

	return strings.Join([]string{"ListExecutionsRequest", string(data)}, " ")
}
