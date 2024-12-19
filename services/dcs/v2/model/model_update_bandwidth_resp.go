package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateBandwidthResp struct {
}

func (o UpdateBandwidthResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBandwidthResp struct{}"
	}

	return strings.Join([]string{"UpdateBandwidthResp", string(data)}, " ")
}
