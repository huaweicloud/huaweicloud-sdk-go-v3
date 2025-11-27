package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LocalSecretReference LocalSecretReference里面包含了访问成员集群的必需凭据的secret信息。例如：- secret.data.token; - secret.data.caBundle。
type LocalSecretReference struct {

	// 资源的命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`
}

func (o LocalSecretReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LocalSecretReference struct{}"
	}

	return strings.Join([]string{"LocalSecretReference", string(data)}, " ")
}
