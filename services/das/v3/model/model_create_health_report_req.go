package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateHealthReportReq struct {

	// 开始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`
}

func (o CreateHealthReportReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthReportReq struct{}"
	}

	return strings.Join([]string{"CreateHealthReportReq", string(data)}, " ")
}
