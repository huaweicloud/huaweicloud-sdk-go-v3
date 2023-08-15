package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConfigMapRequest Request Object
type DeleteConfigMapRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 配置项ID
	ConfigmapId string `json:"configmap_id"`
}

func (o DeleteConfigMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigMapRequest struct{}"
	}

	return strings.Join([]string{"DeleteConfigMapRequest", string(data)}, " ")
}
