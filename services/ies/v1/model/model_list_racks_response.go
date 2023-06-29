package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRacksResponse Response Object
type ListRacksResponse struct {

	// 机柜列表
	Racks *[]Rack `json:"racks,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRacksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRacksResponse struct{}"
	}

	return strings.Join([]string{"ListRacksResponse", string(data)}, " ")
}
