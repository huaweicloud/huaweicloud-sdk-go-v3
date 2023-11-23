package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePluginDraftRequest Request Object
type UpdatePluginDraftRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *PluginDto `json:"body,omitempty"`
}

func (o UpdatePluginDraftRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginDraftRequest struct{}"
	}

	return strings.Join([]string{"UpdatePluginDraftRequest", string(data)}, " ")
}
