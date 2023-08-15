package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEventResponse Response Object
type ShowEventResponse struct {

	// 攻击事件数量
	Total *int32 `json:"total,omitempty"`

	// 攻击事件详情
	Items          *[]ShowEventItems `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventResponse struct{}"
	}

	return strings.Join([]string{"ShowEventResponse", string(data)}, " ")
}
