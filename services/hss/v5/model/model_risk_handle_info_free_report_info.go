package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RiskHandleInfoFreeReportInfo 免费体检报告信息
type RiskHandleInfoFreeReportInfo struct {

	// **参数解释**: 免费体检服务器数量 **取值范围**: 最小值0，最大值2147483647
	FreeCheckHostNum *int64 `json:"free_check_host_num,omitempty"`
}

func (o RiskHandleInfoFreeReportInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskHandleInfoFreeReportInfo struct{}"
	}

	return strings.Join([]string{"RiskHandleInfoFreeReportInfo", string(data)}, " ")
}
