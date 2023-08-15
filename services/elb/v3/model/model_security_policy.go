package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityPolicy 自定义安全策略信息
type SecurityPolicy struct {

	// 自定义安全安全策略的id。
	Id string `json:"id"`

	// 自定义安全策略的项目id。
	ProjectId string `json:"project_id"`

	// 自定义安全策略的名称
	Name string `json:"name"`

	// 自定义安全策略的描述。
	Description string `json:"description"`

	// 自定义安全策略关联的监听器。
	Listeners []ListenerRef `json:"listeners"`

	// 自定义安全策略的TLS协议列表。
	Protocols []string `json:"protocols"`

	// 自定义安全策略的加密套件列表。
	Ciphers []string `json:"ciphers"`

	// 自定义安全策略的创建时间。
	CreatedAt string `json:"created_at"`

	// 自定义安全策略的更新时间。
	UpdatedAt string `json:"updated_at"`
}

func (o SecurityPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityPolicy struct{}"
	}

	return strings.Join([]string{"SecurityPolicy", string(data)}, " ")
}
