package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineBasicDto struct {

	// 流水线id
	Id *int32 `json:"id,omitempty"`

	// 流水线url
	WebUrl *string `json:"web_url,omitempty"`

	// 流水线关联sha值
	Sha *string `json:"sha,omitempty"`

	// 流水线ref
	Ref *string `json:"ref,omitempty"`

	// 流水线状态
	Status *string `json:"status,omitempty"`
}

func (o PipelineBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineBasicDto struct{}"
	}

	return strings.Join([]string{"PipelineBasicDto", string(data)}, " ")
}
