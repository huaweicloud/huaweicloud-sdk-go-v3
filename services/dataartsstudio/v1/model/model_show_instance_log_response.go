package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceLogResponse Response Object
type ShowInstanceLogResponse struct {

	// 是否开启桥接模式
	EnableBridge *bool `json:"enable_bridge,omitempty"`

	// 是否启用配置
	EnableProfile *bool `json:"enable_profile,omitempty"`

	// 是否开启分类
	EnableClassification *bool `json:"enable_classification,omitempty"`

	// 桥接状态
	BridgeStatus *string `json:"bridge_status,omitempty"`

	// 配置状态
	ProfileStatus *string `json:"profile_status,omitempty"`

	// 分类状态
	ClassificationStatus *string `json:"classification_status,omitempty"`

	// 桥接作业日志
	BridgeJobLog *string `json:"bridge_job_log,omitempty"`

	// 配置作业日志
	ProfileJobLog *string `json:"profile_job_log,omitempty"`

	// 分类作业日志
	ClassificationJobLog *string `json:"classification_job_log,omitempty"`
	HttpStatusCode       int     `json:"-"`
}

func (o ShowInstanceLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceLogResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceLogResponse", string(data)}, " ")
}
