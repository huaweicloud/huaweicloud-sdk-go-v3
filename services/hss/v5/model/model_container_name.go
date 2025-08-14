package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerName **参数解释**： 容器实例名称，只有容器类型的告警有 **取值范围**： 字符长度1-256位
type ContainerName struct {
}

func (o ContainerName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerName struct{}"
	}

	return strings.Join([]string{"ContainerName", string(data)}, " ")
}
