package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrainingJobInfo 任务信息
type TrainingJobInfo struct {
	JobType *JobType `json:"job_type,omitempty"`

	// 任务id。
	JobId *string `json:"job_id,omitempty"`

	// 用户id。
	AppUserId *string `json:"app_user_id,omitempty"`

	// 音色名称。该名称会作为资产库中音色模型资产名称。
	VoiceName *string `json:"voice_name,omitempty"`

	// 性别。 * FEMALE: 女性 * MALE: 是男性
	Sex *string `json:"sex,omitempty"`

	// 语言。
	Language *string `json:"language,omitempty"`

	State *JobState `json:"state,omitempty"`

	// 本次任务中该状态出现的次数
	RejectTimes *int32 `json:"reject_times,omitempty"`

	// 当任务状态为成功时呈现,音色模型在资产库中的id。
	AssetId *string `json:"asset_id,omitempty"`

	// 当任务失败时呈现,失败错误码。
	JobFailedCode *string `json:"job_failed_code,omitempty"`

	// 当任务失败时呈现,失败原因。
	JobFailedReason *string `json:"job_failed_reason,omitempty"`

	// 任务创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 任务状态更新时间。
	LastupdateTime *int64 `json:"lastupdate_time,omitempty"`

	// 用户授权书连接。
	VoiceAuthorizationUrl *string `json:"voice_authorization_url,omitempty"`

	// 用户原始输入数据url。
	TrainingInputSourceUrl *string `json:"training_input_source_url,omitempty"`

	CreateType *CreateType `json:"create_type,omitempty"`

	Tag *JobTag `json:"tag,omitempty"`

	// 手机号
	Phone *string `json:"phone,omitempty"`

	// 形象制作任务id
	DhtmsJobId *string `json:"dhtms_job_id,omitempty"`

	// 批次名称
	BatchName *string `json:"batch_name,omitempty"`

	AllocatedResource *VoiceTrainingAllocatedResource `json:"allocated_resource,omitempty"`

	// 模型输出语言类型
	OutputLanguage *string `json:"output_language,omitempty"`

	// 任务是否重做
	IsRemake *bool `json:"is_remake,omitempty"`
}

func (o TrainingJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingJobInfo struct{}"
	}

	return strings.Join([]string{"TrainingJobInfo", string(data)}, " ")
}
