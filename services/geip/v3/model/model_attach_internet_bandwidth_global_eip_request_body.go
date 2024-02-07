package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachInternetBandwidthGlobalEipRequestBody struct {
	DryRun *bool `json:"dry_run,omitempty"`

	GlobalEip *AttachInternetBandwidthGlobalEipRequestBodyGlobalEip `json:"global_eip"`
}

func (o AttachInternetBandwidthGlobalEipRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternetBandwidthGlobalEipRequestBody struct{}"
	}

	return strings.Join([]string{"AttachInternetBandwidthGlobalEipRequestBody", string(data)}, " ")
}
