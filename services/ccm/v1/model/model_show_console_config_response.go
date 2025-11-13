package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsoleConfigResponse Response Object
type ShowConsoleConfigResponse struct {

	// 是否支持企业项目。
	IsSupportEps *bool `json:"is_support_eps,omitempty"`

	// 是否支持SM2算法。
	IsSupportSm2 *bool `json:"is_support_sm2,omitempty"`

	// 是否支持专属加密集群。
	IsSupportDhsm *bool `json:"is_support_dhsm,omitempty"`

	// 专属加密集群支持的region列表。
	DhsmRegions *[]string `json:"dhsm_regions,omitempty"`

	// 是否支持包年包月CA。
	IsSupportYearlyMonthlyCa *bool `json:"is_support_yearly_monthly_ca,omitempty"`

	// 是否支持IAM5鉴权。
	IsSupportIam5 *bool `json:"is_support_iam5,omitempty"`

	// 是否支持OCSP查询。
	IsSupportOcsp  *bool `json:"is_support_ocsp,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowConsoleConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsoleConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowConsoleConfigResponse", string(data)}, " ")
}
