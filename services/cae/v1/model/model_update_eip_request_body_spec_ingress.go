package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEipRequestBodySpecIngress 入网IP信息。
type UpdateEipRequestBodySpecIngress struct {

	// 入网IP带宽。
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 是否关闭入网IP。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o UpdateEipRequestBodySpecIngress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEipRequestBodySpecIngress struct{}"
	}

	return strings.Join([]string{"UpdateEipRequestBodySpecIngress", string(data)}, " ")
}
