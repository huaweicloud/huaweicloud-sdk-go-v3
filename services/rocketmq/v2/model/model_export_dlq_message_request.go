package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportDlqMessageRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *ExportDlqMessageReq `json:"body,omitempty" xml:"body"`
}

func (o ExportDlqMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDlqMessageRequest struct{}"
	}

	return strings.Join([]string{"ExportDlqMessageRequest", string(data)}, " ")
}
