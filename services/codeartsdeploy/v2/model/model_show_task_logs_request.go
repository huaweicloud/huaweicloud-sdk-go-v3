package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskLogsRequest Request Object
type ShowTaskLogsRequest struct {

	// 应用id
	AppId string `json:"app_id"`

	// 部署记录id
	RecordId string `json:"record_id"`

	// 起始位置，默认0；长度1-19
	Offset string `json:"offset"`

	// 结束位置，开始为0；长度1-19
	EndOffset string `json:"end_offset"`

	// 步骤id；如果传，则按step查询日志；长度1-60
	StepId *string `json:"step_id,omitempty"`
}

func (o ShowTaskLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskLogsRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskLogsRequest", string(data)}, " ")
}
