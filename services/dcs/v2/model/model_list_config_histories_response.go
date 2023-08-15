package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigHistoriesResponse Response Object
type ListConfigHistoriesResponse struct {

	// 实例参数修改记录个数
	HistoryNum *int32 `json:"history_num,omitempty"`

	// 实力参数修改记录详情
	Histories      *[]HistoryInfo `json:"histories,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListConfigHistoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigHistoriesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigHistoriesResponse", string(data)}, " ")
}
