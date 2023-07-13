package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStandardTemplateRequest Request Object
type UpdateStandardTemplateRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *StandElementFieldVo `json:"body,omitempty"`
}

func (o UpdateStandardTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStandardTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateStandardTemplateRequest", string(data)}, " ")
}
