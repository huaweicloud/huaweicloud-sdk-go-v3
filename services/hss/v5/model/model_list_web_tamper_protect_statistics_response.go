package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebTamperProtectStatisticsResponse Response Object
type ListWebTamperProtectStatisticsResponse struct {

	// 防护失败的数量
	ProtectFailedNums *int32 `json:"protect_failed_nums,omitempty"`

	// 防护冗余数量
	RedundantNums *int32 `json:"redundant_nums,omitempty"`

	// 未防护数量
	UnprotectNums *int32 `json:"unprotect_nums,omitempty"`

	// 已防护数量
	ProtectedNums *int32 `json:"protected_nums,omitempty"`

	// 防护资产总数
	ProtectTotalNums *int32 `json:"protect_total_nums,omitempty"`

	// 防护事件数量
	ProtectedEventNums *int32 `json:"protected_event_nums,omitempty"`
	HttpStatusCode     int    `json:"-"`
}

func (o ListWebTamperProtectStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebTamperProtectStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListWebTamperProtectStatisticsResponse", string(data)}, " ")
}
