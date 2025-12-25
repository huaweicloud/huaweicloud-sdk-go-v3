package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAnalysisScriptResponse Response Object
type DeleteAnalysisScriptResponse struct {

	// UUID
	ScriptId *string `json:"script_id,omitempty"`

	// 毫秒时间戳
	DeleteTime     *int64 `json:"delete_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteAnalysisScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAnalysisScriptResponse struct{}"
	}

	return strings.Join([]string{"DeleteAnalysisScriptResponse", string(data)}, " ")
}
