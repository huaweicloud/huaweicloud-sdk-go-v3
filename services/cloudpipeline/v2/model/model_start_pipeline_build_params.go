package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartPipelineBuildParams struct {

	// 构建参数名
	Name string `json:"name" xml:"name"`

	// 构建参数值，最大长度为8192
	Value string `json:"value" xml:"value"`
}

func (o StartPipelineBuildParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPipelineBuildParams struct{}"
	}

	return strings.Join([]string{"StartPipelineBuildParams", string(data)}, " ")
}
