package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRepoTagRequestBody struct {

	// 源镜像版本名称
	SourceTag string `json:"source_tag"`

	// 目标镜像版本名称
	DestinationTag string `json:"destination_tag"`

	// 是否覆盖
	Override *bool `json:"override,omitempty"`
}

func (o CreateRepoTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRepoTagRequestBody", string(data)}, " ")
}
