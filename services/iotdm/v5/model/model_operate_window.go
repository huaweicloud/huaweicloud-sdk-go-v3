package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperateWindow 实例维护时间窗，用户在变更实例规格时，可以指定在该时间窗内进行变更。
type OperateWindow struct {

	// **参数说明**：变更时间窗开始时间，UTC时间，格式为：\"HH:mm\"
	StartTime string `json:"start_time"`

	// **参数说明**：变更时间窗结束时间，UTC时间，格式为：\"HH:mm\"
	EndTime string `json:"end_time"`
}

func (o OperateWindow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateWindow struct{}"
	}

	return strings.Join([]string{"OperateWindow", string(data)}, " ")
}
