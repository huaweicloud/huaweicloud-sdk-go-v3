package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSkillBindingsResponse Response Object
type CreateSkillBindingsResponse struct {

	// 批量绑定操作的 TaskFlow 任务 ID。
	JobId *string `json:"job_id,omitempty"`

	// 请求总数。
	Total *int32 `json:"total,omitempty"`

	// 成功数量。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 失败数量。
	FailedCount *int32 `json:"failed_count,omitempty"`

	// 失败详情列表。
	FailedDetails *[]BindingFailedDetail `json:"failed_details,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSkillBindingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSkillBindingsResponse struct{}"
	}

	return strings.Join([]string{"CreateSkillBindingsResponse", string(data)}, " ")
}
