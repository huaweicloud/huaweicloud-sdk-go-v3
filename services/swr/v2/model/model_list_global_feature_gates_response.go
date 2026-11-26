package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalFeatureGatesResponse Response Object
type ListGlobalFeatureGatesResponse struct {

	// 是否支持使用用户的obs桶
	EnableUserDefObs *bool `json:"enableUserDefObs,omitempty"`

	// 是否支持支持企业项目
	EnableEnterprise *bool `json:"enableEnterprise,omitempty"`

	// 是否支持SWR企业版功能
	CerAvailable *bool `json:"cerAvailable,omitempty"`

	// 是否展示创建内网访问的开关按钮
	EnableIntranetAccessSwitch *bool `json:"enableIntranetAccessSwitch,omitempty"`

	// 是否支持Console侧的一键委托
	EnableConsoleAgency *bool `json:"enableConsoleAgency,omitempty"`

	// 是否支持IAM新平面的信任委托
	EnableV5CustomAgency *bool `json:"enableV5CustomAgency,omitempty"`

	// 委托名称
	AgencyName     *string `json:"agencyName,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGlobalFeatureGatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalFeatureGatesResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalFeatureGatesResponse", string(data)}, " ")
}
