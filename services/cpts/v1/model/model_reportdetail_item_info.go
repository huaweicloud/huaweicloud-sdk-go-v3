package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportdetailItemInfo struct {

	// 自定义事务数据
	CustomTransactions *[]string `json:"customTransactions,omitempty" xml:"customTransactions"`

	// aw数据
	DetailDatas *[]DetailDataInfo `json:"detailDatas,omitempty" xml:"detailDatas"`

	Performance *PerformanceInfo `json:"performance,omitempty" xml:"performance"`
}

func (o ReportdetailItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportdetailItemInfo struct{}"
	}

	return strings.Join([]string{"ReportdetailItemInfo", string(data)}, " ")
}
