package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSharedBandwidthRequest struct {

	// 带宽唯一标识  约束： 当前仅支持删除共享带宽
	BandwidthId string `json:"bandwidth_id" xml:"bandwidth_id"`
}

func (o DeleteSharedBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSharedBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DeleteSharedBandwidthRequest", string(data)}, " ")
}
