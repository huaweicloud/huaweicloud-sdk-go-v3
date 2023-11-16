package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginInputsRequest Request Object
type ShowPluginInputsRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *[]PluginPartQueryDto `json:"body,omitempty"`
}

func (o ShowPluginInputsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginInputsRequest struct{}"
	}

	return strings.Join([]string{"ShowPluginInputsRequest", string(data)}, " ")
}
