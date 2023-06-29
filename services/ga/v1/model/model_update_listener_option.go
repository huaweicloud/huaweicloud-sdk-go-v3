package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateListenerOption 更新监听器的详细信息。
type UpdateListenerOption struct {

	// 监听器名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name *string `json:"name,omitempty"`

	// 监听器的描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`

	// 监听端口范围列表。
	PortRanges *[]PortRange `json:"port_ranges,omitempty"`

	ClientAffinity *ClientAffinity `json:"client_affinity,omitempty"`
}

func (o UpdateListenerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerOption struct{}"
	}

	return strings.Join([]string{"UpdateListenerOption", string(data)}, " ")
}
