package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditOperateLogsResponse Response Object
type ListAuditOperateLogsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 操作日志列表
	OperateLog     *[]OperateLogInfo `json:"operate_log,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListAuditOperateLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditOperateLogsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditOperateLogsResponse", string(data)}, " ")
}
