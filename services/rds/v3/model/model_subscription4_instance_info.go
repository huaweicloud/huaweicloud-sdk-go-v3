package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Subscription4InstanceInfo 本地订阅信息。
type Subscription4InstanceInfo struct {

	// 发布服务器来源为云上实例时的发布实例id。
	PublicationInstanceId *string `json:"publication_instance_id,omitempty"`

	// 发布服务器名称。
	PublicationInstanceName *string `json:"publication_instance_name,omitempty"`
}

func (o Subscription4InstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subscription4InstanceInfo struct{}"
	}

	return strings.Join([]string{"Subscription4InstanceInfo", string(data)}, " ")
}
