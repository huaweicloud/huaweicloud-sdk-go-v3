package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateChatResponse Response Object
type UpdateChatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateChatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChatResponse struct{}"
	}

	return strings.Join([]string{"UpdateChatResponse", string(data)}, " ")
}
