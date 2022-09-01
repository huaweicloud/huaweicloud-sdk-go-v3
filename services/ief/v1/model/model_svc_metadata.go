package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 服务详情
type SvcMetadata struct {

	// 自定义标签属性列表
	Labels map[string]string `json:"labels,omitempty" xml:"labels"`

	// 服务名称，只允许英文小写字母、数字、中划线，最大长度64，英文小写字母开头，数字或小写字母结尾
	Name string `json:"name" xml:"name"`
}

func (o SvcMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SvcMetadata struct{}"
	}

	return strings.Join([]string{"SvcMetadata", string(data)}, " ")
}
