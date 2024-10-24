package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessageNotificationPolicyResponse Response Object
type ListMessageNotificationPolicyResponse struct {

	// 符合条件的消息通知策略总数
	Total *int32 `json:"total,omitempty"`

	// 消息通知策略列表信息
	Policies       *[]MessageNotificationPolicy `json:"policies,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListMessageNotificationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageNotificationPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListMessageNotificationPolicyResponse", string(data)}, " ")
}
