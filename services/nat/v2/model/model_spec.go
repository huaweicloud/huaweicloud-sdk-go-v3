package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Spec struct {

	// 规格名称。
	Name *string `json:"name,omitempty"`

	// 规格编号。
	Code *string `json:"code,omitempty"`

	// 规格cbc编码。
	CbcCode *string `json:"cbc_code,omitempty"`

	// 最大规则数。
	RuleMax *int32 `json:"rule_max,omitempty"`

	// 最大连接数。
	SessMax *int32 `json:"sess_max,omitempty"`

	// 最大bps。
	BpsMax *int32 `json:"bps_max,omitempty"`

	// 最大pps。
	PpsMax *int32 `json:"pps_max,omitempty"`

	// 最大qps。
	QpsMax *int32 `json:"qps_max,omitempty"`
}

func (o Spec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Spec struct{}"
	}

	return strings.Join([]string{"Spec", string(data)}, " ")
}
