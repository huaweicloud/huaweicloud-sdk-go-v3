package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTerminalsBindingDesktopsInfoNewResponse Response Object
type ExportTerminalsBindingDesktopsInfoNewResponse struct {

	// 任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportTerminalsBindingDesktopsInfoNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTerminalsBindingDesktopsInfoNewResponse struct{}"
	}

	return strings.Join([]string{"ExportTerminalsBindingDesktopsInfoNewResponse", string(data)}, " ")
}
