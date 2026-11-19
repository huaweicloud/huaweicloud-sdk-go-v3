package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCommonInterfaceResponse Response Object
type ExportCommonInterfaceResponse struct {

	// 导出任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportCommonInterfaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCommonInterfaceResponse struct{}"
	}

	return strings.Join([]string{"ExportCommonInterfaceResponse", string(data)}, " ")
}
