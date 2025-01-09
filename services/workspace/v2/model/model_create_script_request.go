package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScriptRequest Request Object
type CreateScriptRequest struct {
	Body *CreateScriptReq `json:"body,omitempty"`
}

func (o CreateScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScriptRequest struct{}"
	}

	return strings.Join([]string{"CreateScriptRequest", string(data)}, " ")
}
