package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeLoadbalancerResponse Response Object
type UpgradeLoadbalancerResponse struct {

	// 请求ID。 注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	// 升级的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"UpgradeLoadbalancerResponse", string(data)}, " ")
}
