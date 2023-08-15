package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpGroupOption 更新IP地址组的详细信息。
type UpdateIpGroupOption struct {

	// IP地址组名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name *string `json:"name,omitempty"`

	// IP地址组的描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`
}

func (o UpdateIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupOption", string(data)}, " ")
}
