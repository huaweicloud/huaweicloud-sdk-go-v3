package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SvcMetadata 服务详情
type SvcMetadata struct {

	// 自定义标签属性列表
	Labels map[string]string `json:"labels,omitempty"`

	// 服务名称，只允许英文小写字母、数字、中划线，最大长度64，英文小写字母开头，数字或小写字母结尾
	Name string `json:"name"`
}

func (o SvcMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SvcMetadata struct{}"
	}

	return strings.Join([]string{"SvcMetadata", string(data)}, " ")
}
