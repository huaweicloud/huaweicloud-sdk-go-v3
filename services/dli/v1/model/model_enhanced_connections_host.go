package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnhancedConnectionsHost 用户自定义主机信息。
type EnhancedConnectionsHost struct {

	// 自定义主机名称。长度128，数字字母下划线(\"_\")横杠(\"-\")句点(\".\")组成，字母开头。
	Name *string `json:"name,omitempty"`

	// 主机对应的IPv4地址。
	Ip *string `json:"ip,omitempty"`
}

func (o EnhancedConnectionsHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnhancedConnectionsHost struct{}"
	}

	return strings.Join([]string{"EnhancedConnectionsHost", string(data)}, " ")
}
