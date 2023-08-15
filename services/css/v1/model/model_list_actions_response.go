package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListActionsResponse Response Object
type ListActionsResponse struct {

	// 操作记录列表。
	Actions        *[]Actions `json:"actions,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListActionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListActionsResponse struct{}"
	}

	return strings.Join([]string{"ListActionsResponse", string(data)}, " ")
}
