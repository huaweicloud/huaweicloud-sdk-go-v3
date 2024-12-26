package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthInfoExternal 带宽信息
type BandwidthInfoExternal struct {

	// 带宽值
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`

	// 带宽包ID
	GcbId *string `json:"gcb_id,omitempty"`
}

func (o BandwidthInfoExternal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthInfoExternal struct{}"
	}

	return strings.Join([]string{"BandwidthInfoExternal", string(data)}, " ")
}
