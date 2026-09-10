package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteChatResponse Response Object
type DeleteChatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteChatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteChatResponse struct{}"
	}

	return strings.Join([]string{"DeleteChatResponse", string(data)}, " ")
}
