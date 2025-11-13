package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectPublicationMonitorRequest Request Object
type CollectPublicationMonitorRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 发布ID
	PublicationId string `json:"publication_id"`
}

func (o CollectPublicationMonitorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectPublicationMonitorRequest struct{}"
	}

	return strings.Join([]string{"CollectPublicationMonitorRequest", string(data)}, " ")
}
