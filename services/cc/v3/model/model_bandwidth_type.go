package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthType 连接类型。
type BandwidthType struct {
	BandwidthType *BandwidthTypeEnum `json:"bandwidth_type"`
}

func (o BandwidthType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthType struct{}"
	}

	return strings.Join([]string{"BandwidthType", string(data)}, " ")
}
