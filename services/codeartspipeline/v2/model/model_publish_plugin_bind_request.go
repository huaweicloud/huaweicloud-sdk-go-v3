package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishPluginBindRequest Request Object
type PublishPluginBindRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *PublishPluginDto `json:"body,omitempty"`
}

func (o PublishPluginBindRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishPluginBindRequest struct{}"
	}

	return strings.Join([]string{"PublishPluginBindRequest", string(data)}, " ")
}
