package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScriptResponse Response Object
type UpdateScriptResponse struct {

	// 被修改脚本的script_uuid
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScriptResponse struct{}"
	}

	return strings.Join([]string{"UpdateScriptResponse", string(data)}, " ")
}
