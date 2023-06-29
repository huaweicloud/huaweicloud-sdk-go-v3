package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableModelRequest Request Object
type DeleteTableModelRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *IdsParam `json:"body,omitempty"`
}

func (o DeleteTableModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableModelRequest struct{}"
	}

	return strings.Join([]string{"DeleteTableModelRequest", string(data)}, " ")
}
