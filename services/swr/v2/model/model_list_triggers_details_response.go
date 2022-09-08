package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
