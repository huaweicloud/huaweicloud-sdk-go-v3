package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationRequestBodyMetadata 请求数据。
type CreateApplicationRequestBodyMetadata struct {

	// 应用名称。
	Name string `json:"name"`

	// 创建应用附加参数，当前只支持description参数。
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (o CreateApplicationRequestBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationRequestBodyMetadata struct{}"
	}

	return strings.Join([]string{"CreateApplicationRequestBodyMetadata", string(data)}, " ")
}
