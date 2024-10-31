package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeCloudServiceRequest Request Object
type SubscribeCloudServiceRequest struct {

	// iDME服务的类型。  说明：目前仅支持接口开通设计态服务STUDIO  示例：STUDIO
	ServiceType string `json:"service_type"`
}

func (o SubscribeCloudServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeCloudServiceRequest struct{}"
	}

	return strings.Join([]string{"SubscribeCloudServiceRequest", string(data)}, " ")
}
