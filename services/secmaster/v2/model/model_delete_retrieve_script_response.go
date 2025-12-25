package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRetrieveScriptResponse Response Object
type DeleteRetrieveScriptResponse struct {

	// UUID
	ScriptId *string `json:"script_id,omitempty"`

	// 毫秒时间戳
	DeleteTime     *int64 `json:"delete_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteRetrieveScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRetrieveScriptResponse struct{}"
	}

	return strings.Join([]string{"DeleteRetrieveScriptResponse", string(data)}, " ")
}
