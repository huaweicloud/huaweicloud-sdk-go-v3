package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStagePluginsRequest Request Object
type ListStagePluginsRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *StagePluginsQueryDto `json:"body,omitempty"`
}

func (o ListStagePluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStagePluginsRequest struct{}"
	}

	return strings.Join([]string{"ListStagePluginsRequest", string(data)}, " ")
}
