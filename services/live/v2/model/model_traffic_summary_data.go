package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrafficSummaryData struct {

	// 流量，单位为byte。
	Value *int64 `json:"value,omitempty" xml:"value"`

	// 域名。
	Domain *string `json:"domain,omitempty" xml:"domain"`
}

func (o TrafficSummaryData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrafficSummaryData struct{}"
	}

	return strings.Join([]string{"TrafficSummaryData", string(data)}, " ")
}
