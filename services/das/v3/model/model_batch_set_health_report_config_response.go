package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetHealthReportConfigResponse Response Object
type BatchSetHealthReportConfigResponse struct {

	// 各实例设置结果。
	Results        *[]BatchSetResultDto `json:"results,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o BatchSetHealthReportConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetHealthReportConfigResponse struct{}"
	}

	return strings.Join([]string{"BatchSetHealthReportConfigResponse", string(data)}, " ")
}
