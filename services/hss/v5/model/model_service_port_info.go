package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServicePortInfo 容器各服务端口信息
type ServicePortInfo struct {

	// 服务名称
	Desc string `json:"desc"`

	// **参数解释**: 类型 **约束限制**: 不涉及 **取值范围**:   - http：http类端口   - https：https类端口  **默认取值**: 不涉及
	Type *string `json:"type,omitempty"`

	// 默认tcp。可取值集合[tcp，udp]
	Protocol string `json:"protocol"`

	// 用户端口
	UserPort *int32 `json:"user_port,omitempty"`

	// 容器内部端口
	Port int32 `json:"port"`
}

func (o ServicePortInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServicePortInfo struct{}"
	}

	return strings.Join([]string{"ServicePortInfo", string(data)}, " ")
}
