package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyParameterConfigTemplateResponse Response Object
type ModifyParameterConfigTemplateResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyParameterConfigTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyParameterConfigTemplateResponse struct{}"
	}

	return strings.Join([]string{"ModifyParameterConfigTemplateResponse", string(data)}, " ")
}
