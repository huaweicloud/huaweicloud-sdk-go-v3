package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HistoryCommandPage 查询批量分页结构体，定义了分页页码、每页记录数、该页记录的最大Id。
type HistoryCommandPage struct {

	// 本次分页查询结果中最后一条记录的ID，可在下一次分页查询时使用。
	Marker *string `json:"marker,omitempty"`
}

func (o HistoryCommandPage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryCommandPage struct{}"
	}

	return strings.Join([]string{"HistoryCommandPage", string(data)}, " ")
}
