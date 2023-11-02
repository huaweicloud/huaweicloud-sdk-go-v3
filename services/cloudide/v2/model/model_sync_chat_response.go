package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncChatResponse Response Object
type SyncChatResponse struct {

	// result_id
	ResultId *string `json:"result_id,omitempty"`

	// status
	Status *string `json:"status,omitempty"`

	// error message
	ErrorMessage   *string `json:"error_message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SyncChatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncChatResponse struct{}"
	}

	return strings.Join([]string{"SyncChatResponse", string(data)}, " ")
}
