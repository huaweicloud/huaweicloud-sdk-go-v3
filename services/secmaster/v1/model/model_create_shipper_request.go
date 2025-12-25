package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShipperRequest Request Object
type CreateShipperRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateShipperBody `json:"body,omitempty"`
}

func (o CreateShipperRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShipperRequest struct{}"
	}

	return strings.Join([]string{"CreateShipperRequest", string(data)}, " ")
}
