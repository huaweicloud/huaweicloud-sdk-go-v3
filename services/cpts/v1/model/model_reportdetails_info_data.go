package model

import (
	"encoding/json"

	"strings"
)

type ReportdetailsInfoData struct {
	// customTransactions

	CustomTransactions *[]string `json:"customTransactions,omitempty"`
	// detailDatas

	DetailDatas *[]ReportdetailsInfoDetailDatas `json:"detailDatas,omitempty"`
}

func (o ReportdetailsInfoData) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReportdetailsInfoData struct{}"
	}

	return strings.Join([]string{"ReportdetailsInfoData", string(data)}, " ")
}
