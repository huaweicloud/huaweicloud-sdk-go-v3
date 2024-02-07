package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchInternetBandwidthsGlobalEipResponseBody struct {

	// 请求完成的job id
	JobId string `json:"job_id"`
}

func (o BatchInternetBandwidthsGlobalEipResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInternetBandwidthsGlobalEipResponseBody struct{}"
	}

	return strings.Join([]string{"BatchInternetBandwidthsGlobalEipResponseBody", string(data)}, " ")
}
