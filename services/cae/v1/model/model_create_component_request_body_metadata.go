package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentRequestBodyMetadata 请求数据。
type CreateComponentRequestBodyMetadata struct {

	// 组件名称。
	Name string `json:"name"`

	// 创建组件请求体附加参数，当前只支持version参数，此参数必填。
	Annotations map[string]string `json:"annotations"`
}

func (o CreateComponentRequestBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentRequestBodyMetadata struct{}"
	}

	return strings.Join([]string{"CreateComponentRequestBodyMetadata", string(data)}, " ")
}
