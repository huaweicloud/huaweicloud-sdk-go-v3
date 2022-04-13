package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddPipelineJobResponse struct {
	// 管道ID

	PipelineId *string `json:"pipeline_id,omitempty"`
	// 管道错误详情

	CheckInfo      map[string]interface{} `json:"check_info,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o AddPipelineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPipelineJobResponse struct{}"
	}

	return strings.Join([]string{"AddPipelineJobResponse", string(data)}, " ")
}
