package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportbrokensInfo struct {
	BrandBrokens *BrandBrokens `json:"brand_brokens,omitempty" xml:"brand_brokens"`

	// 时间戳
	CommonTimestamps *[]string `json:"commonTimestamps,omitempty" xml:"commonTimestamps"`

	RespcodeBrokens *RespcodeBrokens `json:"respcode_brokens,omitempty" xml:"respcode_brokens"`

	TpsBrokens *TpsBrokens `json:"tps_brokens,omitempty" xml:"tps_brokens"`

	VusersBrokens *VusersBrokens `json:"vusers_brokens,omitempty" xml:"vusers_brokens"`
}

func (o ReportbrokensInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportbrokensInfo struct{}"
	}

	return strings.Join([]string{"ReportbrokensInfo", string(data)}, " ")
}
