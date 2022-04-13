package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResourceHistoryResponse struct {
	// 资源历史列表

	Items *[]HistoryItem `json:"items,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowResourceHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceHistoryResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceHistoryResponse", string(data)}, " ")
}
