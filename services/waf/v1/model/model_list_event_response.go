package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventResponse Response Object
type ListEventResponse struct {

	// 攻击事件数量
	Total *int32 `json:"total,omitempty"`

	// 攻击事件详情
	Items          *[]ListEventItems `json:"items,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventResponse struct{}"
	}

	return strings.Join([]string{"ListEventResponse", string(data)}, " ")
}
