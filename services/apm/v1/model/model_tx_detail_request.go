package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 获取URL跟踪视图详情入参。
type TxDetailRequest struct {

	// 事务名称。
	TxName string `json:"tx_name"`

	// 开始时间。
	StartTime string `json:"start_time"`

	// 结束时间。
	EndTime string `json:"end_time"`
}

func (o TxDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TxDetailRequest struct{}"
	}

	return strings.Join([]string{"TxDetailRequest", string(data)}, " ")
}
