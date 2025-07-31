package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgentDaemonsetDeployTemplateResponse Response Object
type ShowAgentDaemonsetDeployTemplateResponse struct {

	// 是否默认模板
	IsDefault *bool `json:"is_default,omitempty"`

	// 容器运行时配置
	RuntimeInfo *[]RuntimeRequestBody `json:"runtime_info,omitempty"`

	ScheduleInfo   *DeployTemplateInfoScheduleInfo `json:"schedule_info,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowAgentDaemonsetDeployTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentDaemonsetDeployTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowAgentDaemonsetDeployTemplateResponse", string(data)}, " ")
}
