package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishPluginDraftRequest Request Object
type PublishPluginDraftRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *PluginPartQueryDto `json:"body,omitempty"`
}

func (o PublishPluginDraftRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishPluginDraftRequest struct{}"
	}

	return strings.Join([]string{"PublishPluginDraftRequest", string(data)}, " ")
}
