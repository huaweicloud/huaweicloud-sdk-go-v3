package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDlqMessageRequest Request Object
type ExportDlqMessageRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ExportDlqMessageReq `json:"body,omitempty"`
}

func (o ExportDlqMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDlqMessageRequest struct{}"
	}

	return strings.Join([]string{"ExportDlqMessageRequest", string(data)}, " ")
}
