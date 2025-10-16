package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticInterfaceResponse Response Object
type ShowStatisticInterfaceResponse struct {

	// 资源名称
	MetricName *string `json:"metric_name,omitempty"`

	// 接口统计列表
	Metrics        *[]ShowStatisticInterfaceResponseBodyMetrics `json:"metrics,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o ShowStatisticInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticInterfaceResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticInterfaceResponse", string(data)}, " ")
}
