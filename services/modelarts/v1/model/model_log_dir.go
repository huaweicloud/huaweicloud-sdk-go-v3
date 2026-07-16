package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogDir 训练作业可视化日志输出，log_type非空时必填。
type LogDir struct {
	Pfs *PfsSummary `json:"pfs"`
}

func (o LogDir) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogDir struct{}"
	}

	return strings.Join([]string{"LogDir", string(data)}, " ")
}
