package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBasicPluginRequest Request Object
type CreateBasicPluginRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *PluginBasicDto `json:"body,omitempty"`
}

func (o CreateBasicPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBasicPluginRequest struct{}"
	}

	return strings.Join([]string{"CreateBasicPluginRequest", string(data)}, " ")
}
