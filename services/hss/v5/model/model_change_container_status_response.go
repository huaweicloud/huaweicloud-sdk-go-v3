package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeContainerStatusResponse Response Object
type ChangeContainerStatusResponse struct {

	// **参数解释**: 容器ID **取值范围**: 字符长度1-256位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器名称 **取值范围**： 字符长度1-256位
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: 是否成功 **取值范围**: - true：是。 - false：否。
	Success *bool `json:"success,omitempty"`

	// **参数解释**： 状态 **取值范围**： - Running：运行中。 - Terminated：终止。 - Waiting：等待。 - Isolated：已隔离。
	NewStatus      *string `json:"new_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeContainerStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeContainerStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeContainerStatusResponse", string(data)}, " ")
}
