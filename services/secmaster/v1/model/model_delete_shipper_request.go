package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteShipperRequest Request Object
type DeleteShipperRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *DeleteShipperRequestBody `json:"body,omitempty"`
}

func (o DeleteShipperRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShipperRequest struct{}"
	}

	return strings.Join([]string{"DeleteShipperRequest", string(data)}, " ")
}
