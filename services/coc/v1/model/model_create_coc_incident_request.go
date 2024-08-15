package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCocIncidentRequest Request Object
type CreateCocIncidentRequest struct {
	Body *CreateExternalIncidentRequest `json:"body,omitempty"`
}

func (o CreateCocIncidentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCocIncidentRequest struct{}"
	}

	return strings.Join([]string{"CreateCocIncidentRequest", string(data)}, " ")
}
