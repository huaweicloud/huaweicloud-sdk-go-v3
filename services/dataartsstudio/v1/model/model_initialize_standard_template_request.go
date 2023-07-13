package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InitializeStandardTemplateRequest Request Object
type InitializeStandardTemplateRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// action-id=init
	ActionId string `json:"action-id"`

	Body *StandElementFieldVoList `json:"body,omitempty"`
}

func (o InitializeStandardTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitializeStandardTemplateRequest struct{}"
	}

	return strings.Join([]string{"InitializeStandardTemplateRequest", string(data)}, " ")
}
