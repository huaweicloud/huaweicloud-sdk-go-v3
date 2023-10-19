package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseUserBehaviorResponse Response Object
type ParseUserBehaviorResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ParseUserBehaviorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseUserBehaviorResponse struct{}"
	}

	return strings.Join([]string{"ParseUserBehaviorResponse", string(data)}, " ")
}
