package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePluginVersionRequest Request Object
type CreatePluginVersionRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *PluginDto `json:"body,omitempty"`
}

func (o CreatePluginVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginVersionRequest struct{}"
	}

	return strings.Join([]string{"CreatePluginVersionRequest", string(data)}, " ")
}
