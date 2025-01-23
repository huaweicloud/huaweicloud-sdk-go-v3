package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloneLoadbalancerResponse Response Object
type CloneLoadbalancerResponse struct {

	// 新实例相关信息
	LoadbalancerList *[]CloneLoadbalancerResponseBodyLoadbalancerList `json:"loadbalancer_list,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	// 实例复制任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CloneLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"CloneLoadbalancerResponse", string(data)}, " ")
}
