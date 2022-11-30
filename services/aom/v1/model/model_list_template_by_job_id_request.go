package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTemplateByJobIdRequest struct {

	// 作业id。
	JobId string `json:"job_id"`

	Body *ListTemplateByJobIdRequestBody `json:"body,omitempty"`
}

func (o ListTemplateByJobIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateByJobIdRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateByJobIdRequest", string(data)}, " ")
}
