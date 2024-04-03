package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelRemoteAssistanceResponse Response Object
type CancelRemoteAssistanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelRemoteAssistanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelRemoteAssistanceResponse struct{}"
	}

	return strings.Join([]string{"CancelRemoteAssistanceResponse", string(data)}, " ")
}
