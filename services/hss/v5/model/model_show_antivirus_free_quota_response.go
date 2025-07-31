package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAntivirusFreeQuotaResponse Response Object
type ShowAntivirusFreeQuotaResponse struct {

	// 病毒查杀免费扫描次数
	FreeQuota *int32 `json:"free_quota,omitempty"`

	// 当前扫描任务占用的免费扫描次数
	OccupiedFreeQuota *int32 `json:"occupied_free_quota,omitempty"`

	// 病毒查杀界面是否已经赠送过免费次数
	AntivirusAlreadyGiven *bool `json:"antivirus_already_given,omitempty"`

	// 病毒查杀界面应该赠送的免费次数
	AntivirusFreeQuota *int32 `json:"antivirus_free_quota,omitempty"`

	// 月度报告界面是否已经赠送过免费次数
	ReportAlreadyGiven *bool `json:"report_already_given,omitempty"`

	// 月度报告界面应该赠送的免费次数
	ReportFreeQuota *int32 `json:"report_free_quota,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ShowAntivirusFreeQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAntivirusFreeQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowAntivirusFreeQuotaResponse", string(data)}, " ")
}
