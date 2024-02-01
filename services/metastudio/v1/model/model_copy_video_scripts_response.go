package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyVideoScriptsResponse Response Object
type CopyVideoScriptsResponse struct {

	// 新剧本ID
	ScriptId *string `json:"script_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyVideoScriptsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyVideoScriptsResponse struct{}"
	}

	return strings.Join([]string{"CopyVideoScriptsResponse", string(data)}, " ")
}
