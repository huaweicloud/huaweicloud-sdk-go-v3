package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportdetailsInfo struct {
	// data

	Data *[]ReportdetailItemInfo `json:"data,omitempty"`
	// pageIndex

	PageIndex *int32 `json:"pageIndex,omitempty"`
	// pageSize

	PageSize *int32 `json:"pageSize,omitempty"`
	// total

	Total *int32 `json:"total,omitempty"`
}

func (o ReportdetailsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportdetailsInfo struct{}"
	}

	return strings.Join([]string{"ReportdetailsInfo", string(data)}, " ")
}
