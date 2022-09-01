package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流水线信息
type SpecificCommitInfoLastPipeline struct {

	// 流水线id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 提交对应的SHA id
	Sha *string `json:"sha,omitempty" xml:"sha"`

	// 分支名
	Ref *string `json:"ref,omitempty" xml:"ref"`

	// 流水线状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 流水线url
	WebUrl *string `json:"web_url,omitempty" xml:"web_url"`
}

func (o SpecificCommitInfoLastPipeline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecificCommitInfoLastPipeline struct{}"
	}

	return strings.Join([]string{"SpecificCommitInfoLastPipeline", string(data)}, " ")
}
