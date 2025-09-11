package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditOperateLogsNewResponse Response Object
type ListAuditOperateLogsNewResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 操作日志列表
	OperateLog     *[]OperateLogInfo `json:"operate_log,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListAuditOperateLogsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditOperateLogsNewResponse struct{}"
	}

	return strings.Join([]string{"ListAuditOperateLogsNewResponse", string(data)}, " ")
}
