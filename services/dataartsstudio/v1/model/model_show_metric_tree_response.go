package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetricTreeResponse Response Object
type ShowMetricTreeResponse struct {

	// 结构体系
	Architecture   *[]ArchitectureStatistic `json:"architecture,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowMetricTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricTreeResponse struct{}"
	}

	return strings.Join([]string{"ShowMetricTreeResponse", string(data)}, " ")
}
