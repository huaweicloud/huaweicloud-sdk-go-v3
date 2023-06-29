package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentRequestBodyMetadata 请求数据。
type CreateComponentRequestBodyMetadata struct {

	// 组件名称。
	Name string `json:"name"`

	// 组件信息。
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (o CreateComponentRequestBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentRequestBodyMetadata struct{}"
	}

	return strings.Join([]string{"CreateComponentRequestBodyMetadata", string(data)}, " ")
}
