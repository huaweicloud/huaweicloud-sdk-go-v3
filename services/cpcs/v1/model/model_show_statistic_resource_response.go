package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticResourceResponse Response Object
type ShowStatisticResourceResponse struct {

	// 资源名称
	MetricName *string `json:"metric_name,omitempty"`

	// 资源分布
	Datapoints *[]ShowStatisticResourceResponseBodyDatapoints `json:"datapoints,omitempty"`

	// 总数
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowStatisticResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticResourceResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticResourceResponse", string(data)}, " ")
}
