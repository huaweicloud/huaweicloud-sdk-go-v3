package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobLogRequest Request Object
type ShowSparkJobLogRequest struct {

	// 批处理作业的ID。
	BatchId string `json:"batch_id"`

	// 起始日志的行号，默认显示最后100行日志。如果日志不足100行，从0行开始显示。
	From *int32 `json:"from,omitempty"`

	// 当提交的作业进行重试时，会有多个driver日志。index用于指定driver日志的索引号，默认为0。需与type参数一起使用。
	Index *int32 `json:"index,omitempty"`

	// 查询日志的数量。
	Size *int32 `json:"size,omitempty"`

	// 当type填写driver时，输出Spark Driver日志。
	Type *string `json:"type,omitempty"`
}

func (o ShowSparkJobLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobLogRequest struct{}"
	}

	return strings.Join([]string{"ShowSparkJobLogRequest", string(data)}, " ")
}
