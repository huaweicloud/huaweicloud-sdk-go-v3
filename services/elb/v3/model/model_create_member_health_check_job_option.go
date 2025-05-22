package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMemberHealthCheckJobOption 参数解释：创建后端服务器检测任务请求参数。
type CreateMemberHealthCheckJobOption struct {

	// 参数解释：后端服务器所关联的监听器，查询在该监听器下后端服务器的状态。
	ListenerId string `json:"listener_id"`

	// 参数法解释：检查项。  取值范围： - securityGroup：安全组检查。 - networkAcl：子网ACL配置检查。 - config：健康检查端口配置检查。 - all：所有检查项。
	Subject string `json:"subject"`
}

func (o CreateMemberHealthCheckJobOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberHealthCheckJobOption struct{}"
	}

	return strings.Join([]string{"CreateMemberHealthCheckJobOption", string(data)}, " ")
}
