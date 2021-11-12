package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStacksByTagResponse struct {
	Stack *StacksTag `json:"stack,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListStacksByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStacksByTagResponse struct{}"
	}

	return strings.Join([]string{"ListStacksByTagResponse", string(data)}, " ")
}
