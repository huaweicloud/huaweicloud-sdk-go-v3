package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTriggersDetailsResponse Response Object
type ListTriggersDetailsResponse struct {

	// 触发器列表
	Body           *[]Trigger `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListTriggersDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTriggersDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListTriggersDetailsResponse", string(data)}, " ")
}
