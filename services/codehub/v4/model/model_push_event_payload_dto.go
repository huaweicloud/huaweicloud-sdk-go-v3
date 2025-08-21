package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushEventPayloadDto **参数解释：** 推送动态负载。
type PushEventPayloadDto struct {

	// **参数解释：** 提交数量。
	CommitCount *int32 `json:"commit_count,omitempty"`

	// **参数解释：** 操作类型。
	Action *string `json:"action,omitempty"`

	// **参数解释：** ref类型。
	RefType *string `json:"ref_type,omitempty"`

	// **参数解释：** 源Commit。
	CommitFrom *string `json:"commit_from,omitempty"`

	// **参数解释：** 目标Commit。
	CommitTo *string `json:"commit_to,omitempty"`

	// **参数解释：** 分支。
	Ref *string `json:"ref,omitempty"`

	// **参数解释：** 提交标题。
	CommitTitle *string `json:"commit_title,omitempty"`
}

func (o PushEventPayloadDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushEventPayloadDto struct{}"
	}

	return strings.Join([]string{"PushEventPayloadDto", string(data)}, " ")
}
