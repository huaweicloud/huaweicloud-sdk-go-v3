package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscription4InstanceInfo 直接创建订阅时需要的参数。
type CreateSubscription4InstanceInfo struct {

	// 直接创建订阅时，服务器来源为云上实例的发布id。
	PublicationId string `json:"publication_id"`

	// 直接创建订阅时，服务器来源为云上实例的发布名称。
	PublicationName *string `json:"publication_name,omitempty"`
}

func (o CreateSubscription4InstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscription4InstanceInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscription4InstanceInfo", string(data)}, " ")
}
