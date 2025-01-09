package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppInstallJob 应用安装的job。
type AppInstallJob struct {

	// 任务ID。
	Id *string `json:"id,omitempty"`

	// 应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 应用版本。
	AppVersion *string `json:"app_version,omitempty"`

	// 实例ID(桌面或者云应用实例的资源ID)。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例的sid。
	InstanceSid *string `json:"instance_sid,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 操作用户。
	Operator *string `json:"operator,omitempty"`

	// 目标用户。
	Target *string `json:"target,omitempty"`

	JobStatus *JobStatus `json:"job_status,omitempty"`

	// 任务失败时的原因。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o AppInstallJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppInstallJob struct{}"
	}

	return strings.Join([]string{"AppInstallJob", string(data)}, " ")
}
