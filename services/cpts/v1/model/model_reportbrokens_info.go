package model

import (
	"encoding/json"

	"strings"
)

type ReportbrokensInfo struct {
	BrandBrokens *BrandBrokens `json:"brand_brokens,omitempty"`
	// commonTimestamps

	CommonTimestamps *[]string `json:"commonTimestamps,omitempty"`

	RespcodeBrokens *RespcodeBrokens `json:"respcode_brokens,omitempty"`

	TpsBrokens *TpsBrokens `json:"tps_brokens,omitempty"`

	VusersBrokens *VusersBrokens `json:"vusers_brokens,omitempty"`
}

func (o ReportbrokensInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReportbrokensInfo struct{}"
	}

	return strings.Join([]string{"ReportbrokensInfo", string(data)}, " ")
}
