package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistoryDetailsResponse Response Object
type ShowHistoryDetailsResponse struct {

	// 构建任务名称
	JobName *string `json:"job_name,omitempty"`

	// 构建编号
	BuildNumber *int32 `json:"build_number,omitempty"`

	// 构建任务所在项目的ID
	ProjectId *string `json:"project_id,omitempty"`

	// 构建任务所在项目的名称
	ProjectName *string `json:"project_name,omitempty"`

	// 本次构建的参数，Map类型，敏感参数值返回*号
	Parameters map[string]string `json:"parameters,omitempty"`

	// 本次任务的构建步骤详情，返回的步骤为页面可见步骤
	BuildSteps     *[]BuildStep `json:"build_steps,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowHistoryDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoryDetailsResponse", string(data)}, " ")
}
