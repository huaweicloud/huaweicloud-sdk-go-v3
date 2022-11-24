package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建监听器的详细信息。
type CreateListenerOption struct {

	// 监听器名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name string `json:"name"`

	// 监听器的描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`

	Protocol *ListenerProtocol `json:"protocol"`

	// 监听端口范围列表。
	PortRanges []PortRange `json:"port_ranges"`

	ClientAffinity *ClientAffinity `json:"client_affinity,omitempty"`

	// 全球加速实例ID。
	AcceleratorId string `json:"accelerator_id"`

	// 标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o CreateListenerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerOption struct{}"
	}

	return strings.Join([]string{"CreateListenerOption", string(data)}, " ")
}
