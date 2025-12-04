package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeatureGatesResponse Response Object
type ListFeatureGatesResponse struct {

	// 是否开启域名管理
	EnableDomainName *bool `json:"enableDomainName,omitempty"`

	// 老化策略是否支持多条规则
	EnableCombinationRetention *bool `json:"enableCombinationRetention,omitempty"`

	// 是否支持制品扫描
	EnableArtifactScanning *bool `json:"enableArtifactScanning,omitempty"`
	HttpStatusCode         int   `json:"-"`
}

func (o ListFeatureGatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeatureGatesResponse struct{}"
	}

	return strings.Join([]string{"ListFeatureGatesResponse", string(data)}, " ")
}
