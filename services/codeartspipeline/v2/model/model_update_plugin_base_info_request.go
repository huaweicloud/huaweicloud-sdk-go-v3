package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePluginBaseInfoRequest Request Object
type UpdatePluginBaseInfoRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *PluginBasicDto `json:"body,omitempty"`
}

func (o UpdatePluginBaseInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginBaseInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdatePluginBaseInfoRequest", string(data)}, " ")
}
