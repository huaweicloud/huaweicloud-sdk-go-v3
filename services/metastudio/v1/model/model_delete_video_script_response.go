package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVideoScriptResponse Response Object
type DeleteVideoScriptResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteVideoScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVideoScriptResponse struct{}"
	}

	return strings.Join([]string{"DeleteVideoScriptResponse", string(data)}, " ")
}
