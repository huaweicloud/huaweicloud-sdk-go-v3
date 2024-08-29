package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVoiceTrainingJobRequest Request Object
type ListVoiceTrainingJobRequest struct {

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 过滤创建时间<=输入时间的记录。
	CreateUntil *string `json:"create_until,omitempty"`

	// 过滤创建时间>=输入时间的记录。
	CreateSince *string `json:"create_since,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 任务状态，默认所有状态。 可多个状态查询，使用英文逗号分隔。 如state=FAILED,WAITING
	State *string `json:"state,omitempty"`

	// 任务id。
	JobId *string `json:"job_id,omitempty"`

	// 声音名称。
	VoiceName *string `json:"voice_name,omitempty"`

	// 任务标签。
	Tag *string `json:"tag,omitempty"`

	// 训练类型。 * BASIC: 基础版(20句话) * MIDDLE: 进阶版(100句话) * ADVANCE: 高级版 * THIRD_PARTY: 第三方出门问问训练版 * THIRD_PARTY_LJZN: 第三方逻辑智能训练版 * FLEXUS: Flexus版---用的是大模型特征提取
	JobType *string `json:"job_type,omitempty"`
}

func (o ListVoiceTrainingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVoiceTrainingJobRequest struct{}"
	}

	return strings.Join([]string{"ListVoiceTrainingJobRequest", string(data)}, " ")
}
