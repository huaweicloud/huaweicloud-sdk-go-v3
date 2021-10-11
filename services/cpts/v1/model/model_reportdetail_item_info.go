package model

import (
	"encoding/json"

	"strings"
)

type ReportdetailItemInfo struct {
	// customTransactions

	CustomTransactions *[]string `json:"customTransactions,omitempty"`
	// detailDatas

	DetailDatas *[]DetailDataInfo `json:"detailDatas,omitempty"`
}

func (o ReportdetailItemInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReportdetailItemInfo struct{}"
	}

	return strings.Join([]string{"ReportdetailItemInfo", string(data)}, " ")
}
