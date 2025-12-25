package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTtsJobRequest Request Object
type ShowTtsJobRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 过滤创建时间>=输入时间的记录。
	CreateSince *string `json:"create_since,omitempty"`

	// 过滤创建时间<=输入时间的记录。
	CreateUntil *string `json:"create_until,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务类型。 * AUDITION：试听任务 * ASYNC_JOB：异步任务 * WEBSOCKET：websocket接口合成任务
	JobType *string `json:"job_type,omitempty"`

	// tts版本。 * TTS_LLM: 530大模型（V7版本） * TTS_LLM_VC：530大模型VC版本（V7版本） * TTS_LAB：lab小模型（V5版本） * TTS_LAB_GPU：lab小模型GPU版本（V5版本） * GPU_CLONE：V4模型 * TTS_LLM_VQ：VQ模型（V10版本）
	TtsServiceEnum *string `json:"tts_service_enum,omitempty"`

	// 业务类型。
	BusinessType *string `json:"business_type,omitempty"`
}

func (o ShowTtsJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTtsJobRequest struct{}"
	}

	return strings.Join([]string{"ShowTtsJobRequest", string(data)}, " ")
}
