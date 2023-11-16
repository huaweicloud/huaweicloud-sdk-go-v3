package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBasicPluginRequest Request Object
type UpdateBasicPluginRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *PluginBasicDto `json:"body,omitempty"`
}

func (o UpdateBasicPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBasicPluginRequest struct{}"
	}

	return strings.Join([]string{"UpdateBasicPluginRequest", string(data)}, " ")
}
