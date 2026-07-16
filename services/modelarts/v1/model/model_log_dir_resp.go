package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogDirResp 训练作业可视化日志输出，log_type非空时必填。
type LogDirResp struct {
	Pfs *PfsSummaryResp `json:"pfs"`
}

func (o LogDirResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogDirResp struct{}"
	}

	return strings.Join([]string{"LogDirResp", string(data)}, " ")
}
