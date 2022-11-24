package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CopyConfigurationRequest struct {

	// 被复制的参数模板ID。
	ConfigId string `json:"config_id"`

	Body *CopyConfigurationRequestBody `json:"body,omitempty"`
}

func (o CopyConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CopyConfigurationRequest", string(data)}, " ")
}
