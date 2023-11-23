package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePluginDraftRequest Request Object
type CreatePluginDraftRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *PluginDto `json:"body,omitempty"`
}

func (o CreatePluginDraftRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginDraftRequest struct{}"
	}

	return strings.Join([]string{"CreatePluginDraftRequest", string(data)}, " ")
}
