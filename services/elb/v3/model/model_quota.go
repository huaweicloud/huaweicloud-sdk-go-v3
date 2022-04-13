package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 负载均衡相关各类资源的配额信息。仅返回资源的总配额，不包括剩余可用配额。
type Quota struct {
	// 项目ID。

	ProjectId string `json:"project_id"`
	// 负载均衡器配额。 取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。

	Loadbalancer int32 `json:"loadbalancer"`
	// 证书配额。 取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。

	Certificate int32 `json:"certificate"`
	// 监听器配额。 取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。

	Listener int32 `json:"listener"`
	// 转发策略配额。 取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。

	L7policy int32 `json:"l7policy"`
	// 后端云服务器组配额。 取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。

	Pool int32 `json:"pool"`
	// 健康检查配额。 取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。

	Healthmonitor int32 `json:"healthmonitor"`
	// 后端云服务器配额。 取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。

	Member int32 `json:"member"`
	// 单个pool下的member的配额。 取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。

	MembersPerPool int32 `json:"members_per_pool"`
	// IP地址组配额。 取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)

	Ipgroup int32 `json:"ipgroup"`
	// 自定义安全策略配额。 取值： - 大于等于0：表示当前配额数量。 - -1：表示无配额限制。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)

	SecurityPolicy int32 `json:"security_policy"`
}

func (o Quota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}
