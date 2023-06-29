package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSyncPolicyResponse Response Object
type ShowSyncPolicyResponse struct {

	// 策略版本
	PolicyVersion *int64 `json:"policy_version,omitempty"`

	// 策略更新时间
	PolicyUpdateTime *string `json:"policy_updateTime,omitempty"`

	// 权限策略信息
	Policies *[]Policy `json:"policies,omitempty"`

	// 增量权限策略信息
	PolicyDeltas   *[]PolicyDelta `json:"policy_deltas,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowSyncPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSyncPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowSyncPolicyResponse", string(data)}, " ")
}
