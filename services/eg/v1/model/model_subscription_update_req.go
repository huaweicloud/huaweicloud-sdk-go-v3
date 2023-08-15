package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscriptionUpdateReq 更新订阅，全量更新订阅源和目标
type SubscriptionUpdateReq struct {

	// 订阅描述
	Description *string `json:"description,omitempty"`

	// 订阅事件源列表，字段存在则代表全量更新订阅源
	Sources *[]SubscriptionSource `json:"sources,omitempty"`

	// 订阅事件目标列表，字段存在则代表全量更新订阅目标
	Targets *[]SubscriptionTarget `json:"targets,omitempty"`
}

func (o SubscriptionUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionUpdateReq struct{}"
	}

	return strings.Join([]string{"SubscriptionUpdateReq", string(data)}, " ")
}
