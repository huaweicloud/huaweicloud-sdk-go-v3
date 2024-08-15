package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScriptResponse Response Object
type CreateScriptResponse struct {

	// script_uuid
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScriptResponse struct{}"
	}

	return strings.Join([]string{"CreateScriptResponse", string(data)}, " ")
}
