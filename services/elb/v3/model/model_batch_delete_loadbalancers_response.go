package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteLoadbalancersResponse Response Object
type BatchDeleteLoadbalancersResponse struct {

	// 批量删除任务id
	JobId *string `json:"job_id,omitempty"`

	// 待删除的负载均衡器id列表。
	LoadbalancerIds *[]string `json:"loadbalancer_ids,omitempty"`
	HttpStatusCode  int       `json:"-"`
}

func (o BatchDeleteLoadbalancersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLoadbalancersResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteLoadbalancersResponse", string(data)}, " ")
}
