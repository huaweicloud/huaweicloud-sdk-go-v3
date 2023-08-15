package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAcceleratorOption 更新全球加速器的详细信息。
type UpdateAcceleratorOption struct {

	// 全球加速器名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name *string `json:"name,omitempty"`

	// 全球加速器描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`
}

func (o UpdateAcceleratorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAcceleratorOption struct{}"
	}

	return strings.Join([]string{"UpdateAcceleratorOption", string(data)}, " ")
}
