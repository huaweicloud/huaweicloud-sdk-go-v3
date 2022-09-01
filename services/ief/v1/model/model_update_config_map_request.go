package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateConfigMapRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	// 配置项ID
	ConfigmapId string `json:"configmap_id" xml:"configmap_id"`

	Body *UpdateConfigMaps `json:"body,omitempty" xml:"body"`
}

func (o UpdateConfigMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigMapRequest struct{}"
	}

	return strings.Join([]string{"UpdateConfigMapRequest", string(data)}, " ")
}
