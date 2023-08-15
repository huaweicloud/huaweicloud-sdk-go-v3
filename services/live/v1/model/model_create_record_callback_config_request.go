package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecordCallbackConfigRequest Request Object
type CreateRecordCallbackConfigRequest struct {
	Body *RecordCallbackConfigRequest `json:"body,omitempty"`
}

func (o CreateRecordCallbackConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordCallbackConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordCallbackConfigRequest", string(data)}, " ")
}
