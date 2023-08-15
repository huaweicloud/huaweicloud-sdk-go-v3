package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaskRecordRequest Request Object
type DeleteTaskRecordRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o DeleteTaskRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskRecordRequest struct{}"
	}

	return strings.Join([]string{"DeleteTaskRecordRequest", string(data)}, " ")
}
