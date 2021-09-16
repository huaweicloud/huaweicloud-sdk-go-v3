package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSimCardsResponse struct {
	// 每页记录数

	Limit *int64 `json:"limit,omitempty"`
	// 页码

	Offset *int64 `json:"offset,omitempty"`
	// 总数

	Count *int64 `json:"count,omitempty"`
	// sim卡数据集合

	SimCards       *[]SimDeviceVo `json:"sim_cards,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSimCardsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSimCardsResponse struct{}"
	}

	return strings.Join([]string{"ListSimCardsResponse", string(data)}, " ")
}
