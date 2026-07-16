package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceSecurityConfig **参数解释：** 启动业务容器设置信息。
type ServiceSecurityConfig struct {

	// **参数解释：** 启动业务容器时设置的user_id，默认可为空。 **取值范围：** [1, 60000]。
	UserId *int64 `json:"user_id,omitempty"`

	// **参数解释：** 启动业务容器时设置的group_id，默认可为空。 **取值范围：** [1, 60000]。
	GroupId *int64 `json:"group_id,omitempty"`
}

func (o ServiceSecurityConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceSecurityConfig struct{}"
	}

	return strings.Join([]string{"ServiceSecurityConfig", string(data)}, " ")
}
