package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求数据。
type CreateApplicationRequestBodyMetadata struct {

	// 名称。
	Name string `json:"name"`

	// 应用信息。
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (o CreateApplicationRequestBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationRequestBodyMetadata struct{}"
	}

	return strings.Join([]string{"CreateApplicationRequestBodyMetadata", string(data)}, " ")
}
