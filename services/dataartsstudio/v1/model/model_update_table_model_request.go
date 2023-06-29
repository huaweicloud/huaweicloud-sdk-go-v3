package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTableModelRequest Request Object
type UpdateTableModelRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *TableModelUpdateVo `json:"body,omitempty"`
}

func (o UpdateTableModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTableModelRequest struct{}"
	}

	return strings.Join([]string{"UpdateTableModelRequest", string(data)}, " ")
}
