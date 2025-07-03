package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocListHistoryResponseDataV2 历史数据详细
type CocListHistoryResponseDataV2 struct {

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 历史详情
	Info *[]CocTicketHistoryBaseInfoV2 `json:"info,omitempty"`
}

func (o CocListHistoryResponseDataV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocListHistoryResponseDataV2 struct{}"
	}

	return strings.Join([]string{"CocListHistoryResponseDataV2", string(data)}, " ")
}
