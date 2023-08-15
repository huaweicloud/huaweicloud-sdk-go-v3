package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecificCommitInfoLastPipeline 流水线信息
type SpecificCommitInfoLastPipeline struct {

	// 流水线id
	Id *int32 `json:"id,omitempty"`

	// 提交对应的SHA id
	Sha *string `json:"sha,omitempty"`

	// 分支名
	Ref *string `json:"ref,omitempty"`

	// 流水线状态
	Status *string `json:"status,omitempty"`

	// 流水线url
	WebUrl *string `json:"web_url,omitempty"`
}

func (o SpecificCommitInfoLastPipeline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecificCommitInfoLastPipeline struct{}"
	}

	return strings.Join([]string{"SpecificCommitInfoLastPipeline", string(data)}, " ")
}
