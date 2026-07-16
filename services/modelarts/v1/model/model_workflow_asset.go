package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowAsset Workflow资产。
type WorkflowAsset struct {

	// 资产名称。
	Name *string `json:"name,omitempty"`

	// 资产类型，枚举如下: - algorithm：算法 - algorithm2：新算法 - model：模型算法
	Type *string `json:"type,omitempty"`

	// 资产ID，可在AI Gallery中获取。
	ContentId *string `json:"content_id,omitempty"`

	// 订阅ID，可在AI Gallery中获取。
	SubscriptionId *string `json:"subscription_id,omitempty"`

	// 超期时间。
	ExpiredAt *string `json:"expired_at,omitempty"`
}

func (o WorkflowAsset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowAsset struct{}"
	}

	return strings.Join([]string{"WorkflowAsset", string(data)}, " ")
}
