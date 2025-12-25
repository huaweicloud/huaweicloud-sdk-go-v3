package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleShipperAuthorizationRequest Request Object
type HandleShipperAuthorizationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *AuthorizeHandlerRequestBody `json:"body,omitempty"`
}

func (o HandleShipperAuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleShipperAuthorizationRequest struct{}"
	}

	return strings.Join([]string{"HandleShipperAuthorizationRequest", string(data)}, " ")
}
