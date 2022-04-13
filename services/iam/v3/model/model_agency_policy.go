package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AgencyPolicy struct {
	// 权限版本号，创建自定义策略时，该字段值填为“1.1”。 > - 1.0：系统预置的角色。以服务为粒度，提供有限的服务相关角色用于授权。 > - 1.1：策略。IAM最新提供的一种细粒度授权的能力，可以精确到具体服务的操作、资源以及请求条件等。

	Version string `json:"Version"`
	// 授权语句，描述自定义策略的具体内容，不超过8个。

	Statement []AgencyPolicyStatement `json:"Statement"`
}

func (o AgencyPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyPolicy struct{}"
	}

	return strings.Join([]string{"AgencyPolicy", string(data)}, " ")
}
