package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEntitiesResponse Response Object
type ShowEntitiesResponse struct {

	// 技术资产总数
	Count *int32 `json:"count,omitempty"`

	// 技术资产列表
	Entities *[]OpenEntityHeader `json:"entities,omitempty"`

	// scroll_id
	ScrollId       *string `json:"scroll_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowEntitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEntitiesResponse struct{}"
	}

	return strings.Join([]string{"ShowEntitiesResponse", string(data)}, " ")
}
