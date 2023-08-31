package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdviceResult struct {

	// messageId
	MessageId *string `json:"message_id,omitempty"`

	// 状态码
	StatusCode *string `json:"status_code,omitempty"`

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 索引建议
	IndexAdvice *[]IndexAdviceInfo `json:"index_advice,omitempty"`

	// 诊断建议
	TuningAdvice *[]string `json:"tuning_advice,omitempty"`

	// 格式化SQL
	FormattedSql *string `json:"formatted_sql,omitempty"`

	// 原始SQL
	OriginalSql *string `json:"original_sql,omitempty"`

	// 执行计划
	Explain *[]Explain `json:"explain,omitempty"`

	// 表位置信息
	TbPosInfos *[]TbPosInfo `json:"tb_pos_infos,omitempty"`

	FeedbackInfos *FeedbackInfo `json:"feedback_infos,omitempty"`
}

func (o AdviceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdviceResult struct{}"
	}

	return strings.Join([]string{"AdviceResult", string(data)}, " ")
}
