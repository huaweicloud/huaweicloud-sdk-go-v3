package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginOutputsRequest Request Object
type ShowPluginOutputsRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *[]PluginPartQueryDto `json:"body,omitempty"`
}

func (o ShowPluginOutputsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginOutputsRequest struct{}"
	}

	return strings.Join([]string{"ShowPluginOutputsRequest", string(data)}, " ")
}
