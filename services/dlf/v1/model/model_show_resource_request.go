package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResourceRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 资源id.
	ResourceId string `json:"resource_id"`
}

func (o ShowResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceRequest", string(data)}, " ")
}
