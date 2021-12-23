package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

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
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimCardsResponse struct{}"
	}

	return strings.Join([]string{"ListSimCardsResponse", string(data)}, " ")
}
