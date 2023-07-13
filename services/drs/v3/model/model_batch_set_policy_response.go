package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetPolicyResponse Response Object
type BatchSetPolicyResponse struct {

	// 批量设置同步策略响应体
	Results *[]SyncPolicyResp `json:"results,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchSetPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetPolicyResponse struct{}"
	}

	return strings.Join([]string{"BatchSetPolicyResponse", string(data)}, " ")
}
