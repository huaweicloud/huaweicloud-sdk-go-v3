package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnhancedConnectionHost 用户自定义主机信息。
type EnhancedConnectionHost struct {

	// 自定义主机名称。长度128，数字字母下划线(\"_\")横杠(\"-\")句点(\".\")组成，字母开头。
	Name *string `json:"name,omitempty"`

	// 主机对应的IPv4地址。
	Ip *string `json:"ip,omitempty"`
}

func (o EnhancedConnectionHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnhancedConnectionHost struct{}"
	}

	return strings.Join([]string{"EnhancedConnectionHost", string(data)}, " ")
}
