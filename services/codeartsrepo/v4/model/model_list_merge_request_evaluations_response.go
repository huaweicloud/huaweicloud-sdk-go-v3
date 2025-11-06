package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestEvaluationsResponse Response Object
type ListMergeRequestEvaluationsResponse struct {

	// **参数解释：** 评价id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 合并请求id。
	MergeRequestId *int32 `json:"merge_request_id,omitempty"`

	// **参数解释：** 分数。
	Level *int32 `json:"level,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 评价文本内容。
	Content *string `json:"content,omitempty"`

	User *UserBasicDto `json:"user,omitempty"`

	// 自定义评价
	CustomEvaluations *[]CustomEvaluationDto `json:"custom_evaluations,omitempty"`
	HttpStatusCode    int                    `json:"-"`
}

func (o ListMergeRequestEvaluationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestEvaluationsResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestEvaluationsResponse", string(data)}, " ")
}
