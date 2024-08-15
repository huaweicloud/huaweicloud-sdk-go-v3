package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetScriptResponse Response Object
type GetScriptResponse struct {
	Data           *ScriptDetailModel `json:"data,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o GetScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetScriptResponse struct{}"
	}

	return strings.Join([]string{"GetScriptResponse", string(data)}, " ")
}
