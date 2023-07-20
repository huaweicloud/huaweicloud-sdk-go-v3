package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobMonitorInfoReq 查询作业监控信息
type ShowJobMonitorInfoReq struct {

	// 作业ID列表。
	JobIds []int64 `json:"job_ids"`
}

func (o ShowJobMonitorInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobMonitorInfoReq struct{}"
	}

	return strings.Join([]string{"ShowJobMonitorInfoReq", string(data)}, " ")
}
