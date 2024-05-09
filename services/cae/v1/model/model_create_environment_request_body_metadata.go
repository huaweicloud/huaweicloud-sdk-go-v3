package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnvironmentRequestBodyMetadata 请求数据。
type CreateEnvironmentRequestBodyMetadata struct {

	// 创建环境请求体附加参数。 - vpc_id：创建环境绑定的VPC的ID。 - group_name：创建环境绑定的SWR组织的组织名称。 - type：环境类型，当前仅支持exclusive类型。 - subnet_id: 创建环境绑定的VPC子网的ID。 - security_group_id：创建环境绑定的安全组的ID，可不填，不填由CAE后台自动创建。
	Annotations map[string]string `json:"annotations,omitempty"`

	// 环境名称。
	Name string `json:"name"`
}

func (o CreateEnvironmentRequestBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentRequestBodyMetadata struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentRequestBodyMetadata", string(data)}, " ")
}
