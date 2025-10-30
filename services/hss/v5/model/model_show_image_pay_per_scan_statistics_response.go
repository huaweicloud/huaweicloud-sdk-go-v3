package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImagePayPerScanStatisticsResponse Response Object
type ShowImagePayPerScanStatisticsResponse struct {

	// 仓库镜像扫描成功次数
	RepositoryScanNum *int32 `json:"repository_scan_num,omitempty"`

	// CICD镜像扫描成功次数
	CicdScanNum *int32 `json:"cicd_scan_num,omitempty"`

	// 免费扫描次数
	FreeQuotaNum *int32 `json:"free_quota_num,omitempty"`

	// 是否赠送过免费次数
	AlreadyGiven *bool `json:"already_given,omitempty"`

	// 登录容器镜像界面赠送的次数
	ImageFreeQuota *int32 `json:"image_free_quota,omitempty"`

	HighRisk *ImageTypeRiskInfo `json:"high_risk,omitempty"`

	HasRisk *ImageTypeRiskInfo `json:"has_risk,omitempty"`

	Total *ImageTypeRiskInfo `json:"total,omitempty"`

	Unscan         *ImageTypeRiskInfo `json:"unscan,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowImagePayPerScanStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImagePayPerScanStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowImagePayPerScanStatisticsResponse", string(data)}, " ")
}
