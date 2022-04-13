package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateConfigMapRequest struct {
	// 铂金版实例ID，专业版实例为空值

	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *ConfigMaps `json:"body,omitempty"`
}

func (o CreateConfigMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigMapRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigMapRequest", string(data)}, " ")
}
