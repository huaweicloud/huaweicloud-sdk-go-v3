package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportdetailItemInfo struct {
	// customTransactions

	CustomTransactions *[]string `json:"customTransactions,omitempty"`
	// detailDatas

	DetailDatas *[]DetailDataInfo `json:"detailDatas,omitempty"`
}

func (o ReportdetailItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportdetailItemInfo struct{}"
	}

	return strings.Join([]string{"ReportdetailItemInfo", string(data)}, " ")
}
