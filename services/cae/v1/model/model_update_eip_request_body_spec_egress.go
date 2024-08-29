package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEipRequestBodySpecEgress 出网IP信息。
type UpdateEipRequestBodySpecEgress struct {

	// 出网IP带宽。
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 是否关闭出网IP。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o UpdateEipRequestBodySpecEgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEipRequestBodySpecEgress struct{}"
	}

	return strings.Join([]string{"UpdateEipRequestBodySpecEgress", string(data)}, " ")
}
