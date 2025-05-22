package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreLoadbalancerResponse Response Object
type RestoreLoadbalancerResponse struct {

	// 请求ID。 注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	// ELB实例的类型，v2或者v3。
	Type *string `json:"type,omitempty"`

	// 还原负载均衡器的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestoreLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"RestoreLoadbalancerResponse", string(data)}, " ")
}
