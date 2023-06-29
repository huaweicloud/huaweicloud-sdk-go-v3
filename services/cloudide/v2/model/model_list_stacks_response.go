package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStacksResponse Response Object
type ListStacksResponse struct {
	Stack *StacksTags `json:"stack,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListStacksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStacksResponse struct{}"
	}

	return strings.Join([]string{"ListStacksResponse", string(data)}, " ")
}
