package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishPluginRequest Request Object
type PublishPluginRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *PublishPluginDto `json:"body,omitempty"`
}

func (o PublishPluginRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishPluginRequest struct{}"
	}

	return strings.Join([]string{"PublishPluginRequest", string(data)}, " ")
}
