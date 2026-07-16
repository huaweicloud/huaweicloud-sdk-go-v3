package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolRuntimeMetricsRequest Request Object
type ShowPoolRuntimeMetricsRequest struct {
}

func (o ShowPoolRuntimeMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolRuntimeMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowPoolRuntimeMetricsRequest", string(data)}, " ")
}
