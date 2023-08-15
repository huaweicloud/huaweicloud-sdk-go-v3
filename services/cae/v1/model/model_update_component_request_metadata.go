package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateComponentRequestMetadata struct {

	// 组件名称。
	Name string `json:"name"`

	// 更新组件请求体附加参数，当前只支持version参数，此参数必填。
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (o UpdateComponentRequestMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentRequestMetadata struct{}"
	}

	return strings.Join([]string{"UpdateComponentRequestMetadata", string(data)}, " ")
}
