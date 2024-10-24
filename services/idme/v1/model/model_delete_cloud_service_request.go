package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudServiceRequest Request Object
type DeleteCloudServiceRequest struct {

	// iDME服务的类型。  说明：目前仅支持删除CLOUD_LINK按需资源  示例：CLOUD_LINKX
	ServiceType string `json:"service_type"`

	// 待删除的实例ID。
	InstanceId string `json:"instance_id"`
}

func (o DeleteCloudServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudServiceRequest struct{}"
	}

	return strings.Join([]string{"DeleteCloudServiceRequest", string(data)}, " ")
}
