package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddPipelineJobResponse struct {

	// 管道ID
	PipelineId *string `json:"pipeline_id,omitempty" xml:"pipeline_id"`

	// 管道错误详情
	CheckInfo      map[string]interface{} `json:"check_info,omitempty" xml:"check_info"`
	HttpStatusCode int                    `json:"-"`
}

func (o AddPipelineJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddPipelineJobResponse struct{}"
	}

	return strings.Join([]string{"AddPipelineJobResponse", string(data)}, " ")
}
