package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStandardTemplateRequest Request Object
type CreateStandardTemplateRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *StandElementFieldVo `json:"body,omitempty"`
}

func (o CreateStandardTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStandardTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateStandardTemplateRequest", string(data)}, " ")
}
