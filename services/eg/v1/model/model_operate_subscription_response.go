package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type OperateSubscriptionResponse struct {

	// 操作失败的订阅个数
	FailedCount *int32 `json:"failed_count,omitempty" xml:"failed_count"`

	// 失败信息
	Events         *[]SubscriptionOperateRespEvents `json:"events,omitempty" xml:"events"`
	HttpStatusCode int                              `json:"-"`
}

func (o OperateSubscriptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperateSubscriptionResponse struct{}"
	}

	return strings.Join([]string{"OperateSubscriptionResponse", string(data)}, " ")
}
